package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	//"github.com/evrone/go-clean-template/internal/entity"
	"yarus-tz/internal/entity"
	"yarus-tz/internal/usecase"

	"github.com/evrone/go-clean-template/pkg/logger"
)

type exchangeRoutes struct {
	e usecase.ExchangeUseCase
	l logger.Interface
}

func newExchangeRoutes(handler *gin.RouterGroup, e usecase.ExchangeUseCase, l logger.Interface) {
	r := &exchangeRoutes{e, l}

	h := handler.Group("/exchange")
	{
		h.GET("/:currency", r.doExchange)
	}
}

// @Summary     Exchange
// @Description Translate a text
// @ID          do-exchange
// @Tags  	    exchangee
// @Accept      json
// @Produce     json
// @Param       request body doExchangeRequest true "Set up Exchange"
// @Success     200 {object} entity.Exchange
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /exchange/{currency} [get]
func (r *exchangeRoutes) doExchange(c *gin.Context) {

	currency := c.Param("currency")
	if currency == "" {
		//TODO get random value from map
		panic("implement me")
	}

	exchange, err := r.e.Exchange(
		c.Request.Context(),
		entity.Exchange{
			Currency: currency,
		},
	)
	if err != nil {
		r.l.Error(err, "http - v1 - doTranslate")
		//errorResponse(c, http.StatusInternalServerError, "exchange service problems")

		return
	}

	c.JSON(http.StatusOK, gin.H{
		exchange.Currency: exchange.Response,
	})
}
