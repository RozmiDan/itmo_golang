package main

import (
	"os"

	"github.com/RozmiDan/url_shortener/internal/config"
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

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware_logger.MyLogger(logger))
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

}
