package main

import (
	"context"
	"fmt"
)

type metricSeverice struct {
	next PriceFetcher
}

func NewMetricService(next PriceFetcher) PriceFetcher {
	return &metricSeverice{
		next: next,
	}
}

func (s *metricSeverice) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	fmt.Println("Pushing metrics to prometheus")
	// your metrics storage. Push to prometheus (gauge, counters)
	return s.next.FetchPrice(ctx, ticker)
}
