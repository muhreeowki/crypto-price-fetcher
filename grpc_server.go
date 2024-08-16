package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"

	"github.com/muhreeowki/price-fetcher/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCPriceFetcherServer struct {
	proto.UnimplementedPriceFetcherServer
	svc PriceFetcher
}

func (s *GRPCPriceFetcherServer) FetchPrice(ctx context.Context, req *proto.PriceRequest) (*proto.PriceResponse, error) {
	ctx = context.WithValue(ctx, "requestID", rand.Intn(10000000))
	price, err := s.svc.FetchPrice(ctx, req.Ticker)
	if err != nil {
		return nil, err
	}
	return &proto.PriceResponse{
		Price:  &price,
		Ticker: req.Ticker,
	}, err
}

func NewGRPCPriceFetcher(svc PriceFetcher) *GRPCPriceFetcherServer {
	return &GRPCPriceFetcherServer{
		svc: svc,
	}
}

func MakeAndRunGRPCServer(listenAddr string, svc PriceFetcher) error {
	grpcPriceFetcher := NewGRPCPriceFetcher(svc)

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	reflection.Register(server)
	proto.RegisterPriceFetcherServer(server, grpcPriceFetcher)

	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	fmt.Printf("gRPC listening on %v\n", listenAddr)

	return server.Serve(listener)
}
