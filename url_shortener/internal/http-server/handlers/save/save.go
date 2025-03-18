package save_handler

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/RozmiDan/url_shortener/internal/storage"
	"github.com/RozmiDan/url_shortener/internal/usecase/random"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"
)

const aliasLength = 6

type URLSaver interface {
	SaveURL(urlToSave string, alias string) (int64, error)
}

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Alias  string `json:"alias,omitempty"`
}

func NewSaveHandler(logger *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.save.newsavehandler"

		logger = logger.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			logger.Error("failed to decode request body")
			render.JSON(w, r, Response{
				Status: "Error",
				Error:  "failed to decode request",
			})
			return
		}

		logger.Info("request body decoded", slog.Any("request", req))

		validate := validator.New()
		if err := validate.Struct(req); err != nil {
			logger.Error("validation error", slog.Any("error", err))
			render.JSON(w, r, Response{
				Status: "Error",
				Error:  "invalid request parameters",
			})
			return
		}

		alias := req.Alias
		if alias == "" {
			alias = random.NewAliasForURL(aliasLength)
		}

		_, err = urlSaver.SaveURL(req.URL, alias)
		if err != nil {
			if errors.Is(err, storage.ErrURLExists) {
				logger.Error("URL already exists", slog.String("URL", req.URL))
				render.JSON(w, r, Response{
					Status: "Error",
					Error:  "URL already exists",
				})
				return
			}
			logger.Error("failed to save URL", slog.String("error", err.Error()))
			render.JSON(w, r, Response{
				Status: "Error",
				Error:  "failed to save URL",
			})
			return
		}
		render.JSON(w, r, Response{
			Status: "OK",
			Alias:  alias,
		})
	}
}
