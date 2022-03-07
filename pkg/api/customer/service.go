package customer

import (
	"context"
	"simple-api-gateway/pkg/api/customer/model"
)

type Service interface {
	List(ctx context.Context) ([]model.Customer, error)
	Get(ctx context.Context, request *model.GetRequest) (*model.Customer, error)
	Create(ctx context.Context, request *model.CreateRequest) (int64, error)
	Update(ctx context.Context, request *model.UpdateRequest) error
	Delete(ctx context.Context, request *model.DeleteRequest) error
}
