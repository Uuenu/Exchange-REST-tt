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

// type historyResponse struct {
// 	History []entity.Translation `json:"history"`
// }

// @Summary     Show history
// @Description Show all translation history
// @ID          history
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Success     200 {object} historyResponse
// @Failure     500 {object} response
// @Router      /translation/history [get]

// func (r *translationRoutes) history(c *gin.Context) {
// 	translations, err := r.t.History(c.Request.Context())
// 	if err != nil {
// 		r.l.Error(err, "http - v1 - history")
// 		//errorResponse(c, http.StatusInternalServerError, "database problems")

// 		return
// 	}

// 	c.JSON(http.StatusOK, nil)
// }

// @Summary     Translate
// @Description Translate a text
// @ID          do-translate
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Param       request body doTranslateRequest true "Set up translation"
// @Success     200 {object} entity.Translation
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /translation/do-translate [post]
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
