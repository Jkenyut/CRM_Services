package repository_customer

import (
	"context"
	"crm_service/app/model/model_customer"
)

type InterfaceRepoCustomer interface {
	CreateCustomer(ctx context.Context, req model_customer.RequestCustomer) (int, error)
	GetCustomerByEmail(ctx context.Context, req model_customer.RequestCustomerEmail) (status int, err error, res model_customer.Customer)
	GetCustomerById(ctx context.Context, req uint64) (status int, err error, res model_customer.Customer)
	GetCountRowsCustomer(ctx context.Context) (status int, err error, res model_customer.Customer)
	GetAllCustomer(ctx context.Context, page uint64, limit uint64, firstname string, lastname string) (status int, err error, res []model_customer.Customer)
	UpdateCustomerById(ctx context.Context, id uint64, req model_customer.RequestUpdateCustomer) (int, error)
	DeleteCustomerById(ctx context.Context, id uint64) (int, error)
}
