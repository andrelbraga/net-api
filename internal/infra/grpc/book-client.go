package grpc

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewBookClient
// host: net-grpc -> docker-compose,
// host: localhost:5001 -> go run cmd/main.go
func NewBookClient() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Print(err.Error())
	}
	return conn, nil
}
