package connection

import (
	"google.golang.org/grpc"
	"log"
	"simple-api-gateway/config"
	"sync"
)

var customerOnce sync.Once
var customerGrpcClient *grpc.ClientConn

func GetCustomerGrpcClient() *grpc.ClientConn {
	customerOnce.Do(func() {
		customerGrpcClient = newCustomerGrpcClient()
	})
	return customerGrpcClient
}

func newCustomerGrpcClient() *grpc.ClientConn {
	conn, err := grpc.Dial(config.Env.CustomerGrpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return conn
}
