package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"yarus-tz/internal/entity"
	"yarus-tz/internal/usecase"

	"yarus-tz/pkg/logger"
)

type exchangeRoutes struct {
	e usecase.ExchangeUseCase
	l logger.Interface
}

func newExchangeRoutes(handler *gin.RouterGroup, e usecase.ExchangeUseCase, l logger.Interface) {
	r := &exchangeRoutes{e, l}

	h := handler.Group("/exchange").Use(tokenMiddleware(l, e))
	{
		h.GET("/:currency", r.doExchange)
		h.GET("", r.doExchange)
	}
}

type doExchangeRequest struct {
}

// @Summary     Exchange
// @Description Exchange currency on Ruble
// @ID          do-exchange
// @Tags  	    exchangee
// @Accept      json
// @Produce     json
// @Param       request body doExchangeRequest true "Set up Exchange"
// @Success     200 {object} response
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /exchange [get]
func (r *exchangeRoutes) doExchange(c *gin.Context) {

	currency := c.Param("currency")

	exchange, err := r.e.Exchange(
		c.Request.Context(),
		entity.Exchange{
			Currency: currency,
		},
	)
	if err != nil {
		r.l.Error(err, "http - v1 - doTranslate")
		errorResponse(c, http.StatusInternalServerError, "exchange service problems")

		return
	}

	c.JSON(http.StatusOK, gin.H{
		exchange.Currency: exchange.Response,
	})
}
