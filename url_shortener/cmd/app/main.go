package main

import (
	"net/http"
	"os"

	"github.com/RozmiDan/url_shortener/internal/config"
	save_handler "github.com/RozmiDan/url_shortener/internal/http-server/handlers"
	middleware_logger "github.com/RozmiDan/url_shortener/internal/http-server/middleware"
	"github.com/RozmiDan/url_shortener/internal/storage/sqlite"
	"github.com/RozmiDan/url_shortener/pkg/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	cnfg := config.MustLoad()
	logger := logger.NewLogger(cnfg.Env)

	logger.Info("url-shortner started")
	logger.Debug("debug mode")

	storage, err := sqlite.New(cnfg.StoragePath)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	_ = storage

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware_logger.MyLogger(logger))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Post("/url", save_handler.NewSaveHandler(logger, storage))

	logger.Info("starting server")

	server := http.Server{
		Addr:         cnfg.HttpInfo.Port,
		Handler:      router,
		ReadTimeout:  cnfg.HttpInfo.Timeout,
		WriteTimeout: cnfg.HttpInfo.Timeout,
		IdleTimeout:  cnfg.HttpInfo.IdleTimeout,
	}

	err = server.ListenAndServe()
	if err != nil {
		logger.Error("Server error", err)
	}

	logger.Info("Finishing programm")
}
