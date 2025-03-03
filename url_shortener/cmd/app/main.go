package main

import (
	"fmt"

	"github.com/RozmiDan/url_shortener/internal/config"
	"github.com/RozmiDan/url_shortener/pkg/logger"
)

func main() {
	cnfg := config.MustLoad();
	logger := logger.NewLogger(cnfg.Env)
	fmt.Println("helo")
	logger.Debug("Hello")
	fmt.Println(cnfg)
}