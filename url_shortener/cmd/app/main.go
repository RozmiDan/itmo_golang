package main

import (
	"fmt"
	"os"

	"github.com/RozmiDan/url_shortener/internal/config"
	"github.com/RozmiDan/url_shortener/internal/storage/sqlite"
	"github.com/RozmiDan/url_shortener/pkg/logger"
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
	res, err := storage.GetURL("google")
	if err != nil {
		logger.Error(err.Error())
	}
	fmt.Println(res)

}
