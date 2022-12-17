package usecase

import (
	"context"
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

	return e, nil
}
