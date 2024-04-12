package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/adrianowsh/exchange-rate-api/internal/dto"
)

type ExchangeRate struct {
	Code       string `json:"code"`
	CodeIn     string `json:"code_in"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `jnson:"var_bid"`
	PctChange  string `json:"pct_change"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"tinmestamp"`
	CreateDate string `json:"create_date"`
}

const cotacao_file = "cotacao.txt"
const url_exchange_rate = "http://localhost:8080/cotacao?currency=USD-BRL"
const context_time_duration = 300 * time.Millisecond

func main() {
	ratech := make(chan dto.Response)

	ctx, cancel := context.WithTimeout(context.Background(), context_time_duration)
	defer cancel()

	go func(ch chan<- dto.Response) {
		resp, err := getResultExchangeRateApi(ctx)
		if err != nil {
			ch <- dto.Response{Data: "", Err: err}
		}

		ch <- dto.Response{Data: resp, Err: nil}
	}(ratech)

	select {
	case resp := <-ratech:
		if resp.Err != nil {
			log.Fatalln(resp.Err)
		} else {
			writeBidInformation(resp.Data)
		}
	case <-ctx.Done():
		log.Fatalln("context timeout")
	}
}

func getResultExchangeRateApi(ctx context.Context) (string, error) {
	resp, err := getExchangeRateApi(ctx)
	if err != nil {
		return "", err
	}

	jsonResult, _ := json.Marshal(resp)
	return string(jsonResult), nil
}

func getExchangeRateApi(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url_exchange_rate, nil)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	body, _ := io.ReadAll(res.Body)

	var exchangesRates []ExchangeRate
	err = json.Unmarshal([]byte(body), &exchangesRates)
	if err != nil {
		return "'", err
	}

	return exchangesRates[0].Bid, nil
}

func writeBidInformation(bid string) {
	file, err := os.Create(cotacao_file)
	if err != nil {
		panic(fmt.Sprintf("error on create txt: %v", err))
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("BID: %s", bid))

	if err != nil {
		panic(fmt.Sprintf("error on write txt: %v", err))
	}

	println("file create successfully")
}
