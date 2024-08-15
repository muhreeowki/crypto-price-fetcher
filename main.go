package main

import (
	"flag"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "listendAddress the service should run on in the format of host:port")
	flag.Parse()

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	server := NewJSONAPIServer(*listenAddr, svc)
	server.Start()
}
