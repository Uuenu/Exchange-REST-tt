package v1

import (
	"fmt"
	"net/http"
	"yarus-tz/internal/usecase"
	"yarus-tz/pkg/logger"

	"github.com/gin-gonic/gin"
)

func tokenMiddleware(l logger.Interface, e usecase.ExchangeUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.Request.Header
		var key string

		for k, v := range header["X-Api-Key"] {
			fmt.Println(k, v)
			key = v
		}

		if key != "123321" && len(header["X-Api-Key"]) != 1 {
			l.Info("http - v1 - exchange - middleware tokenMiddleware - Invalid X-API-KEY")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Next()
	}
}
