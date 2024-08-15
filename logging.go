package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

// loggingService is a service that logs the time taken to fetch a price and the result of the fetch.
type loggingService struct {
	next PriceFetcher
}

// FetchPrice is a method on loggingService that fetches the price of a ticker and logs the time taken and result.
func (s *loggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"requestID": ctx.Value("requestID"),
			"took":      time.Since(begin),
			"err":       err,
			"price":     price,
		}).Info("fetchPrice")
	}(time.Now())

	return s.next.FetchPrice(ctx, ticker)
}

// NewLoggingService returns a new logging service with the given price fetcher service
func NewLoggingService(next PriceFetcher) PriceFetcher {
	return &loggingService{
		next: next,
	}
}
