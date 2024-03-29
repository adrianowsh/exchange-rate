package abstractions

import (
	"context"

	"github.com/adrianowsh/exchange-rate-api/internal/dto"
	"github.com/adrianowsh/exchange-rate-api/internal/infra/db"
)

type ExchangeRateRepositoryInterface interface {
	Create(ctx context.Context, entity dto.UsdBrlDTO) error
	List(ctx context.Context) ([]db.ExchangesRate, error)
}
