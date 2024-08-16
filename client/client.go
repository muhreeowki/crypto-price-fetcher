package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/muhreeowki/price-fetcher/proto"
	"github.com/muhreeowki/price-fetcher/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client is a client that fetches the price of a ticker from a remote service.
type Client struct {
	endpoint string
}

// New returns a new client with the given endpoint.
func New(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}

func NewGRPCPriceFetcherClient(serverAddr string) (proto.PriceFetcherClient, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.NewClient(serverAddr, opts...)
	if err != nil {
		return nil, err
	}

	grpcClient := proto.NewPriceFetcherClient(conn)

	return grpcClient, nil
}

// FetchPrice fetches the price of a ticker from the provided endpoint.
func (c *Client) FetchPrice(ctx context.Context, ticker string) (*types.PriceResponse, error) {
	endpoint := fmt.Sprintf("%s?ticker=%s", c.endpoint, ticker)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		httpErr := map[string]any{}
		if err := json.NewDecoder(res.Body).Decode(&httpErr); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("service responded with non 200 status code: %s", httpErr["error"])
	}

	priceResp := new(types.PriceResponse)
	if err := json.NewDecoder(res.Body).Decode(&priceResp); err != nil {
		return nil, err
	}

	return priceResp, nil
}
