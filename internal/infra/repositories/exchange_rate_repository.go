package repositories

import (
	"context"
	"database/sql"

	"github.com/adrianowsh/exchange-rate-api/internal/dto"
	"github.com/adrianowsh/exchange-rate-api/internal/infra/db"
)

type ExchangeRateRepository struct {
	DB *sql.DB
}

func NewExchangeRateRepository(db *sql.DB) *ExchangeRateRepository {
	return &ExchangeRateRepository{DB: db}
}

func (repo *ExchangeRateRepository) Create(ctx context.Context, dto dto.UsdBrlDTO) error {
	queries := db.New(repo.DB)

	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	qtx := queries.WithTx(tx)

	err = qtx.CreateExchangeRate(ctx, db.CreateExchangeRateParams{
		Code:       sql.NullString{String: dto.Code, Valid: true},
		CodeUn:     sql.NullString{String: dto.CodeIn, Valid: true},
		Name:       sql.NullString{String: dto.Name, Valid: true},
		High:       sql.NullString{String: dto.High, Valid: true},
		Low:        sql.NullString{String: dto.Low, Valid: true},
		VarBid:     sql.NullString{String: dto.VarBid, Valid: true},
		PctChange:  sql.NullString{String: dto.PctChange, Valid: true},
		Bid:        sql.NullString{String: dto.Bid, Valid: true},
		Ask:        sql.NullString{String: dto.Ask, Valid: true},
		Timestamp:  sql.NullString{String: dto.Timestamp, Valid: true},
		CreateDate: sql.NullString{String: dto.CreateDate, Valid: true},
	})

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (repo *ExchangeRateRepository) List(ctx context.Context) ([]db.ExchangesRate, error) {
	queries := db.New(repo.DB)

	results, err := queries.ListExchangeRate(ctx)
	if err != nil {
		return nil, err
	}

	return results, nil
}
