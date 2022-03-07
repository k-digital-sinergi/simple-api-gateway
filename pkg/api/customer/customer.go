package customer

import (
	"context"
	"encoding/json"
	connection "simple-api-gateway/connection/grpc"
	"simple-api-gateway/pkg/api/customer/model"
	proto "simple-api-gateway/pkg/api/customer/proto"
)

type Customer struct{}

func New() *Customer {
	return &Customer{}
}

func (c *Customer) List(ctx context.Context) ([]model.Customer, error) {
	client := proto.NewCustomerServiceClient(connection.GetCustomerGrpcClient())
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

func (c *Customer) Create(ctx context.Context, request *model.CreateRequest) (int64, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return 0, err
	}

	client := proto.NewCustomerServiceClient(connection.GetCustomerGrpcClient())
	response, err := client.Create(ctx, &proto.Request{Body: body})
	if err != nil {
		return 0, err
	}

	var custID int64
	if err = json.Unmarshal(response.Body, &custID); err != nil {
		return 0, err
	}

	return custID, nil
}

func (c *Customer) Update(ctx context.Context, request *model.UpdateRequest) error {
	body, err := json.Marshal(request)
	if err != nil {
		return err
	}

	client := proto.NewCustomerServiceClient(connection.GetCustomerGrpcClient())
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

	client := proto.NewCustomerServiceClient(connection.GetCustomerGrpcClient())
	_, err = client.Delete(ctx, &proto.Request{Body: body})
	if err != nil {
		return err
	}

	return nil
}
