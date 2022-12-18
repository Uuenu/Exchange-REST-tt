package usecase

import (
	"context"
	"fmt"
	"yarus-tz/internal/entity"
	"yarus-tz/pkg/logger"
)

// ExchangeUseCase -.
type ExchangeUseCase struct {
	log *logger.Logger
}

func NewExchangeUseCase(l *logger.Logger) *ExchangeUseCase {
	return &ExchangeUseCase{
		log: l,
	}
}

func (uc *ExchangeUseCase) Exchange(ctx context.Context, e entity.Exchange) (entity.Exchange, error) {

	var valutes entity.DailyValues
	valutes.SetValutes()

	var v entity.Valute
	if e.Currency == "" {
		v = valutes.PickRandomValute()

	} else {
		v = valutes.Valutes[e.Currency]
	}

	empty := entity.Valute{}
	if v == empty {
		//TODO added logger and json error
		return entity.Exchange{}, fmt.Errorf("ExchangeUseCase - Exchagne - invalid currency")
	}
	e.Currency = v.CharCode
	e.Response = fmt.Sprintf("%d %s равен %f рублям", v.Nominal, v.Name, v.Value)

	return e, nil
}
