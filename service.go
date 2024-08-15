package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// PriceFetcher is an interface that can fetch a price.
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// priceFetcher implements the PriceFetcher interface.
type priceFetcher struct{}

type coinBaseData struct {
	Amount   string
	Base     string
	Currency string
}

type coinBaseResponse struct {
	Data coinBaseData
}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	// Important to not use types in business logic, do not use json representations
	return FetchCoinBaseAPI(ctx, ticker)
}

// FetchCoinBaseAPI is a function to fetch the price of a ticker from coinbase.
func FetchCoinBaseAPI(ctx context.Context, ticker string) (float64, error) {
	url := fmt.Sprintf("https://api.coinbase.com/v2/prices/%v-USD/buy", ticker)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	// Error handling
	switch res.StatusCode {
	case http.StatusBadRequest:
		return 0, fmt.Errorf("invalid request format")
	case http.StatusUnauthorized:
		return 0, fmt.Errorf("you do not have access to the requested resource")
	case http.StatusNotFound:
		return 0, fmt.Errorf("ticker not found")
	}
	// Marshal the response into a struct
	defer res.Body.Close()
	var coinBaseRes coinBaseResponse
	err = json.NewDecoder(res.Body).Decode(&coinBaseRes)
	if err != nil {
		return 0, err
	}

	price, err := strconv.ParseFloat(coinBaseRes.Data.Amount, 64)
	if err != nil {
		return 0, err
	}

	return price, nil
}
