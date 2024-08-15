package main

import (
	"context"
	"fmt"
)

// metricService is a service that pushes metrics to prometheus
type metricSeverice struct {
	next PriceFetcher
}

// NewMetricService returns a new metric service with the given price fetcher service
func NewMetricService(next PriceFetcher) PriceFetcher {
	return &metricSeverice{
		next: next,
	}
}

// FetchPrice is a method on metricService that fetches the price of a ticker.
func (s *metricSeverice) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	fmt.Println("Pushing metrics to prometheus")
	// your metrics storage. Push to prometheus (gauge, counters)
	return s.next.FetchPrice(ctx, ticker)
}
