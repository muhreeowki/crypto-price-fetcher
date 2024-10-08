package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/muhreeowki/price-fetcher/types"
)

type APIFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type JSONAPIServer struct {
	svc        PriceFetcher
	listenAddr string
}

// NewJSONAPIServer returns a new JSONAPIServer with the given listenAddr and price fetcher service
func NewJSONAPIServer(listenAddr string, svc PriceFetcher) *JSONAPIServer {
	return &JSONAPIServer{
		listenAddr: listenAddr,
		svc:        svc,
	}
}

// Run starts the JSON API server and listens on the given listenAddr
func (s *JSONAPIServer) Run() {
	http.HandleFunc("/", makeHTTPHandlerFunc(s.handleFetchPrice))

	fmt.Printf("json API listening on %v\n", s.listenAddr)
	http.ListenAndServe(s.listenAddr, nil)
}

// handleFetchPrice is a handler function that fetches the price for a given ticker from the price fetcher service
func (s *JSONAPIServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")

	price, err := s.svc.FetchPrice(ctx, ticker)
	if err != nil {
		return err
	}

	priceResponse := types.PriceResponse{
		Price:  price,
		Ticker: ticker,
	}

	return writeJSON(w, http.StatusOK, &priceResponse)
}

// makeHTTPHandlerFunc is a function that takes an APIFunc and returns an http.HandlerFunc
func makeHTTPHandlerFunc(apiFn APIFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", rand.Intn(10000000))

	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFn(ctx, w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

// writeJSON is a function to write a JSON object to the response writer
func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
