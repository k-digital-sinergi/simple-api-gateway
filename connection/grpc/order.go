package connection

import (
	"google.golang.org/grpc"
	"log"
	"simple-api-gateway/config"
	"sync"
)

var orderOnce *sync.Once
var orderGrpcClient *grpc.ClientConn

func GetOrderGrpcClient() *grpc.ClientConn {
	orderOnce.Do(func() {
		orderGrpcClient = newOrderGrpcClient()
	})
	return orderGrpcClient
}

func newOrderGrpcClient() *grpc.ClientConn {
	conn, err := grpc.Dial(config.Env.OrderGrpcAddr)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}
