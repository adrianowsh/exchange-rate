package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
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
	ctx, cancel := context.WithTimeout(context.Background(), context_time_duration)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url_exchange_rate, nil)
	if err != nil {
		panic(fmt.Sprintf("error on the request => %v", err))
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(fmt.Sprintf("error on read all body => %v", err))
	}

	body, _ := io.ReadAll(res.Body)

	var exchangesRates []ExchangeRate
	err = json.Unmarshal([]byte(body), &exchangesRates)
	if err != nil {
		panic(fmt.Sprintf("error on serialized to struct => %v", err))
	}

	file, err := os.Create(cotacao_file)
	if err != nil {
		panic(fmt.Sprintf("error on create txt => %v", err))
	}

	defer file.Close()

	for _, exchangeRate := range exchangesRates {
		writeBidInformation(exchangeRate.Bid, file)
	}
}

func writeBidInformation(bid string, file *os.File) {
	_, err := file.WriteString(fmt.Sprintf("BID: %s", bid))
	fmt.Println("File created successfully")
	if err != nil {
		panic(fmt.Sprintf("error on write txt => %v", err))
	}
}
