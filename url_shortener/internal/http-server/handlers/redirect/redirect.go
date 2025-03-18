package redirect_handler

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/RozmiDan/url_shortener/internal/storage"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type URLGetter interface {
	GetURL(alias string) (string, error)
}

type Response struct {
	err    error
	status string
}

func RedirectHandlerConstructor(logger *slog.Logger, urlGetter URLGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const info = "redirect_handler.RedirectHandlerConstruction"
		logger = logger.With(slog.String("op", info),
			slog.String("request_id", middleware.GetReqID(r.Context())))

		reqAlias := chi.URLParam(r, "alias")

		if reqAlias == "" {
			logger.Error("requested alias is empty")
			render.JSON(w, r, Response{
				err:    errors.New("requested alias is empty"),
				status: "Error",
			})
			return
		}

		logger.Info("request alias is valid")

		url, err := urlGetter.GetURL(reqAlias)

		if err != nil {
			if errors.Is(err, storage.ErrURLNotFound) {
				logger.Error("URL not found", slog.String("alias", reqAlias))
				render.JSON(w, r, Response{
					err:    err,
					status: "Error",
				})
				return
			}
			logger.Error("Error while getting url", slog.String("alias", reqAlias))
			render.JSON(w, r, Response{
				err:    err,
				status: "Error",
			})
			return
		}
		logger.Info("url was found", slog.String("url", url))

		http.Redirect(w, r, url, http.StatusFound)
	}
}
