package app

import (
	"yarus-tz/config"
	v1 "yarus-tz/internal/controller/http/v1"
	"yarus-tz/internal/usecase"
	"yarus-tz/pkg/logger"

	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	//Valute Rest
	//valute.RunValuteRest()

	//Use case
	exchangeUseCase := usecase.NewExchangeUseCase(l)

	//HTTP Server
	handler := gin.Default()
	v1.NewRouter(handler, l, *exchangeUseCase)
	//handler.Run(cfg.HTTP.Port)
	handler.Run()
}
