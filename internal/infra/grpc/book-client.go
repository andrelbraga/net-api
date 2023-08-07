package grpc

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewBookClient
func NewBookClient() (*grpc.ClientConn, error) {
	/* host: net-grpc -> docker-compose, host: localhost:5001 -> go run cmd/main.go */
	conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Print(err.Error())
	}
	return conn, nil
}
