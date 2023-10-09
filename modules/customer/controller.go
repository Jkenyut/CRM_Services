package customer

import (
	"context"
	"crm_service/entity"
	"crm_service/model"
	"math"
	"net/http"
)

type CustomerControllerInterface interface {
	CreateCustomer(ctx context.Context, req RequestCustomer) (entity.DefaultResponse, int, entity.DefaultResponse)
	GetCustomerById(ctx context.Context, id uint64) (entity.DefaultResponse, int, entity.DefaultResponse)
	GetAllCustomer(ctx context.Context, page uint64, req RequestGetAllCustomer) (entity.DefaultResponse, int, entity.DefaultResponse)
	UpdateCustomerById(ctx context.Context, id uint64, req RequestUpdateCustomer) (entity.DefaultResponse, int, entity.DefaultResponse)
	DeleteCustomerById(ctx context.Context, id uint64) (entity.DefaultResponse, int, entity.DefaultResponse)
}

type customerControllerStruct struct {
	customerRepository CustomerRepoInterface
}

func (c customerControllerStruct) CreateCustomer(ctx context.Context, req RequestCustomer) (entity.DefaultResponse, int, entity.DefaultResponse) {
	//var customerCustomer model.Customer
	var response entity.DefaultResponse
	var errorMessage entity.DefaultResponse

	// create acustomer
	status, err := c.customerRepository.CreateCustomer(ctx, req)
	if err != nil {
		errorMessage = entity.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	response = entity.DefaultSuccessResponseWithMessage("customer created", status, "success")
	return response, status, errorMessage
}

func (c customerControllerStruct) GetCustomerById(ctx context.Context, id uint64) (entity.DefaultResponse, int, entity.DefaultResponse) {
	var customerRepo model.Customer
	var response entity.DefaultResponse
	var errorMessage entity.DefaultResponse

	//get data by id
	status, err := c.customerRepository.GetCustomerById(ctx, id, &customerRepo)
	if err != nil {
		errorMessage = entity.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}
	response = entity.DefaultSuccessResponseWithMessage("Get customer", status, customerRepo)
	return response, status, errorMessage
}

func (c customerControllerStruct) GetAllCustomer(ctx context.Context, page uint64, req RequestGetAllCustomer) (entity.DefaultResponse, int, entity.DefaultResponse) {
	var customerRepo []model.Customer
	var customerCountRepo model.Customer
	var response entity.DefaultResponse
	var errorMessage entity.DefaultResponse

	var limit uint64 = 30
	status, err := c.customerRepository.GetCountRowsCustomer(ctx, &customerCountRepo)
	if err != nil {
		errorMessage = entity.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	status, err = c.customerRepository.GetAllCustomer(ctx, page, limit, req, &customerRepo)
	if err != nil {
		errorMessage = entity.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	resMessage := FindAllCustomer{
		Page:       page,
		PerPage:    uint64(len(customerRepo)),
		TotalPages: uint64(math.Ceil(float64(customerCountRepo.Total) / float64(limit))),
		Data:       customerRepo,
	}

	response = entity.DefaultSuccessResponseWithMessage("Get all customer", status, resMessage)

	return response, status, errorMessage
}

func (c customerControllerStruct) UpdateCustomerById(ctx context.Context, id uint64, req RequestUpdateCustomer) (entity.DefaultResponse, int, entity.DefaultResponse) {
	var customerRepo model.Customer
	var response entity.DefaultResponse
	var errorMessage entity.DefaultResponse

	//check authorization
	if id == 1 {
		errorMessage = entity.DefaultErrorResponseWithMessage("not authorization update", http.StatusUnauthorized)
		return response, http.StatusUnauthorized, errorMessage
	}

	//update data by id
	status, err := c.customerRepository.UpdateCustomerById(ctx, id, req)
	if err != nil {
		errorMessage = entity.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	//repo
	status, err = c.customerRepository.GetCustomerById(ctx, id, &customerRepo)
	if err != nil {
		errorMessage = entity.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	response = entity.DefaultSuccessResponseWithMessage("Get customer", status, customerRepo)
	return response, status, errorMessage
}

func (c customerControllerStruct) DeleteCustomerById(ctx context.Context, id uint64) (entity.DefaultResponse, int, entity.DefaultResponse) {
	var response entity.DefaultResponse
	var errorMessage entity.DefaultResponse
	//check authorization
	if id == 1 {
		errorMessage = entity.DefaultErrorResponseWithMessage("not authorization delete", http.StatusUnauthorized)
		return response, http.StatusUnauthorized, errorMessage
	}

	//repo
	status, err := c.customerRepository.DeleteCustomerById(ctx, id)
	if err != nil {
		errorMessage = entity.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}
	response = entity.DefaultSuccessResponseWithMessage("delete customer", status, "true")
	return response, status, errorMessage
}
