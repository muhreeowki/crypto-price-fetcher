package main

import (
	"flag"
)

func main() {
	var (
		jsonListenAddr = flag.String("json", ":3000", "listen address of the json transport in the format of host:port")
		grpcListenAddr = flag.String("grpc", ":4000", "listen address of the grpc transport in the format of host:port")
		// ctx            = context.Background()
		svc = NewLoggingService(NewMetricService(&priceFetcher{}))
	)
	flag.Parse()

	go MakeAndRunGRPCServer(*grpcListenAddr, svc)

	jsonServer := NewJSONAPIServer(*jsonListenAddr, svc)
	jsonServer.Run()
}
