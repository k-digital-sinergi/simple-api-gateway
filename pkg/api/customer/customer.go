package customer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"simple-api-gateway/pkg/api/customer/model"
	proto "simple-api-gateway/pkg/api/customer/proto"
	"simple-api-gateway/redis"
	"simple-api-gateway/rpc"
	"strconv"
	"time"
)

const (
	keyPrefix = "customer:"
)

type Customer struct{}

func New() *Customer {
	return &Customer{}
}

func (c *Customer) List(ctx context.Context) ([]model.Customer, error) {
	client := proto.NewCustomerServiceClient(rpc.GetCustomerGrpcClient())
	response, err := client.List(ctx, &proto.Empty{})
	if err != nil {
		return nil, err
	}

	var customers []model.Customer
	if err = json.Unmarshal(response.Body, &customers); err != nil {
		return nil, err
	}

	return customers, nil
}

func (c *Customer) Get(ctx context.Context, request *model.GetRequest) (*model.Customer, error) {
	conn := redis.GetRedisConnection(ctx)
	key := fmt.Sprintf("%s%s", keyPrefix, strconv.FormatInt(request.CustID, 10))
	data, err := redis.Get(ctx, conn, key)
	if err != nil {
		return nil, err
	}
	if data != "" {
		log.Println("get from redis")
		var customer model.Customer
		err := json.Unmarshal([]byte(data), &customer)
		if err != nil {
			return nil, err
		}
		return &customer, nil
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	client := proto.NewCustomerServiceClient(rpc.GetCustomerGrpcClient())
	response, err := client.Get(ctx, &proto.Request{Body: body})
	if err != nil {
		return nil, err
	}

	var customer model.Customer
	if err = json.Unmarshal(response.Body, &customer); err != nil {
		return nil, err
	}

	err = redis.Set(ctx, conn, key, string(response.Body), 1*time.Minute)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (c *Customer) Create(ctx context.Context, request *model.CreateRequest) (int64, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return 0, err
	}

	client := proto.NewCustomerServiceClient(rpc.GetCustomerGrpcClient())
	response, err := client.Create(ctx, &proto.Request{Body: body})
	if err != nil {
		return 0, err
	}

	var custID int64
	if err = json.Unmarshal(response.Body, &custID); err != nil {
		return 0, err
	}

	key := fmt.Sprintf("%s%s", keyPrefix, strconv.FormatInt(custID, 10))
	err = redis.Del(ctx, redis.GetRedisConnection(ctx), key)
	if err != nil {
		return 0, err
	}

	return custID, nil
}

func (c *Customer) Update(ctx context.Context, request *model.UpdateRequest) error {
	body, err := json.Marshal(request)
	if err != nil {
		return err
	}

	client := proto.NewCustomerServiceClient(rpc.GetCustomerGrpcClient())
	_, err = client.Update(ctx, &proto.Request{Body: body})
	if err != nil {
		return err
	}

	return nil
}

func (c *Customer) Delete(ctx context.Context, request *model.DeleteRequest) error {
	body, err := json.Marshal(request)
	if err != nil {
		return err
	}

	client := proto.NewCustomerServiceClient(rpc.GetCustomerGrpcClient())
	_, err = client.Delete(ctx, &proto.Request{Body: body})
	if err != nil {
		return err
	}

	return nil
}
