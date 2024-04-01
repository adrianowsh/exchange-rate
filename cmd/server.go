package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/adrianowsh/exchange-rate-api/internal/dto"
	"github.com/adrianowsh/exchange-rate-api/internal/infra/db"
	"github.com/adrianowsh/exchange-rate-api/internal/infra/repositories"
	_ "github.com/mattn/go-sqlite3"
)

const context_time_duration = 200 * time.Millisecond
const database_name = "file:sqlite-database.db"

func main() {

	http.HandleFunc("/cotacao", SearchExchangeRateHandler)
	http.HandleFunc("/all", GetRegisteredExchangeRateHandler)
	http.ListenAndServe(":8080", nil)
}

func GetRegisteredExchangeRateHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), context_time_duration)
	defer cancel()

	db, err := sql.Open("sqlite3", database_name)
	if err != nil {
		panic(fmt.Sprintf("error on database access ==> %v", err))
	}
	defer db.Close()

	if r.URL.Path != "/all" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	dto, err := getAll(ctx, db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Contant-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto)
}

func SearchExchangeRateHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), context_time_duration)
	defer cancel()

	db, err := sql.Open("sqlite3", database_name)
	if err != nil {
		panic(fmt.Sprintf("error on database access ==> %v", err))
	}
	defer db.Close()

	if r.URL.Path != "/cotacao" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	currencyParam := r.URL.Query().Get("currency")
	if currencyParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dto, err := searchExchangeRateAndSave(ctx, db, currencyParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Contant-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto)
}

func searchExchangeRateAndSave(ctx context.Context, db *sql.DB, currency string) (*[]dto.UsdBrlDTO, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://economia.awesomeapi.com.br/json/daily/%s", currency), nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var results []dto.UsdBrlDTO
	err = json.Unmarshal([]byte(body), &results)
	if err != nil {
		return nil, err
	}

	repo := repositories.NewExchangeRateRepository(db)
	for _, result := range results {
		saveExchangeRate(repo, result)
	}
	return &results, nil
}

func getAll(ctx context.Context, db *sql.DB) (*[]db.ExchangesRate, error) {
	repo := repositories.NewExchangeRateRepository(db)
	exchangeRates, err := repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return &exchangeRates, nil
}

func saveExchangeRate(repo *repositories.ExchangeRateRepository, dto dto.UsdBrlDTO) error {
	err := repo.Create(dto)
	if err != nil {
		return err
	}
	return nil
}
