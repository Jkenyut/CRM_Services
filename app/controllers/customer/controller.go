package customer

import (
	"context"
	"crm_service/app/model"
	"crm_service/app/model/original"
	"math"
	"net/http"
)

type CustomerControllerInterface interface {
	CreateCustomer(ctx context.Context, req RequestCustomer) (original.DefaultResponse, int, original.DefaultResponse)
	GetCustomerById(ctx context.Context, id uint64) (original.DefaultResponse, int, original.DefaultResponse)
	GetAllCustomer(ctx context.Context, page uint64, req RequestGetAllCustomer) (original.DefaultResponse, int, original.DefaultResponse)
	UpdateCustomerById(ctx context.Context, id uint64, req RequestUpdateCustomer) (original.DefaultResponse, int, original.DefaultResponse)
	DeleteCustomerById(ctx context.Context, id uint64) (original.DefaultResponse, int, original.DefaultResponse)
}

type customerControllerStruct struct {
	customerRepository CustomerRepoInterface
}

func (c customerControllerStruct) CreateCustomer(ctx context.Context, req RequestCustomer) (original.DefaultResponse, int, original.DefaultResponse) {
	//var customerCustomer model.Customer
	var response original.DefaultResponse
	var errorMessage original.DefaultResponse

	// create acustomer
	status, err := c.customerRepository.CreateCustomer(ctx, req)
	if err != nil {
		errorMessage = original.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	response = original.DefaultSuccessResponseWithMessage("customer created", status, "success")
	return response, status, errorMessage
}

func (c customerControllerStruct) GetCustomerById(ctx context.Context, id uint64) (original.DefaultResponse, int, original.DefaultResponse) {
	var customerRepo model.Customer
	var response original.DefaultResponse
	var errorMessage original.DefaultResponse

	//get data by id
	status, err := c.customerRepository.GetCustomerById(ctx, id, &customerRepo)
	if err != nil {
		errorMessage = original.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}
	response = original.DefaultSuccessResponseWithMessage("Get customer", status, customerRepo)
	return response, status, errorMessage
}

func (c customerControllerStruct) GetAllCustomer(ctx context.Context, page uint64, req RequestGetAllCustomer) (original.DefaultResponse, int, original.DefaultResponse) {
	var customerRepo []model.Customer
	var customerCountRepo model.Customer
	var response original.DefaultResponse
	var errorMessage original.DefaultResponse

	var limit uint64 = 30
	status, err := c.customerRepository.GetCountRowsCustomer(ctx, &customerCountRepo)
	if err != nil {
		errorMessage = original.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	status, err = c.customerRepository.GetAllCustomer(ctx, page, limit, req, &customerRepo)
	if err != nil {
		errorMessage = original.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	resMessage := FindAllCustomer{
		Page:       page,
		PerPage:    uint64(len(customerRepo)),
		TotalPages: uint64(math.Ceil(float64(customerCountRepo.Total) / float64(limit))),
		Data:       customerRepo,
	}

	response = original.DefaultSuccessResponseWithMessage("Get all customer", status, resMessage)

	return response, status, errorMessage
}

func (c customerControllerStruct) UpdateCustomerById(ctx context.Context, id uint64, req RequestUpdateCustomer) (original.DefaultResponse, int, original.DefaultResponse) {
	var customerRepo model.Customer
	var response original.DefaultResponse
	var errorMessage original.DefaultResponse

	//check authorization
	if id == 1 {
		errorMessage = original.DefaultErrorResponseWithMessage("not authorization update", http.StatusUnauthorized)
		return response, http.StatusUnauthorized, errorMessage
	}

	//update data by id
	status, err := c.customerRepository.UpdateCustomerById(ctx, id, req)
	if err != nil {
		errorMessage = original.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	//repo
	status, err = c.customerRepository.GetCustomerById(ctx, id, &customerRepo)
	if err != nil {
		errorMessage = original.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	response = original.DefaultSuccessResponseWithMessage("Get customer", status, customerRepo)
	return response, status, errorMessage
}

func (c customerControllerStruct) DeleteCustomerById(ctx context.Context, id uint64) (original.DefaultResponse, int, original.DefaultResponse) {
	var response original.DefaultResponse
	var errorMessage original.DefaultResponse
	//check authorization
	if id == 1 {
		errorMessage = original.DefaultErrorResponseWithMessage("not authorization delete", http.StatusUnauthorized)
		return response, http.StatusUnauthorized, errorMessage
	}

	//repo
	status, err := c.customerRepository.DeleteCustomerById(ctx, id)
	if err != nil {
		errorMessage = original.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}
	response = original.DefaultSuccessResponseWithMessage("delete customer", status, "true")
	return response, status, errorMessage
}
