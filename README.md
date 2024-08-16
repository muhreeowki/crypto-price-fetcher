# Crypto Price Fetcher microservice

This is a repo with code for a crypto price fetcher microservice I built in Go to learn how to architect microservices.
The service runs a server that fetches crypto prices from various exchanges and returns them to the client. It uses both **_gRPC_** and **_REST_** API architectures for communication between the client and the server. This is so that the client can choose the method of communication that best suits their needs.
Its built fully with the go standard library (except for logrus for logging).

## Requirements

- Go 1.22
- make
- docker

## Installation

### Install Protobuf

```bash
sudo apt install -y protobuf-compiler protoc
```

### Install Protobuf and gRPC Dependencies

```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

NOTE: Make sure your go bin directory is in your PATH

```bash
PATH="$PATH:$HOME/go/bin"
```

## Usage

To start the server run the following command

```bash
make run
```
