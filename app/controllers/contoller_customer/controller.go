package contoller_customer

import (
	"context"
	"crm_service/app/model"
	"crm_service/app/model/origin"
	"math"
	"net/http"
)

type CustomerControllerInterface interface {
	CreateCustomer(ctx context.Context, req RequestCustomer) (origin.DefaultResponse, int, origin.DefaultResponse)
	GetCustomerById(ctx context.Context, id uint64) (origin.DefaultResponse, int, origin.DefaultResponse)
	GetAllCustomer(ctx context.Context, page uint64, req RequestGetAllCustomer) (origin.DefaultResponse, int, origin.DefaultResponse)
	UpdateCustomerById(ctx context.Context, id uint64, req RequestUpdateCustomer) (origin.DefaultResponse, int, origin.DefaultResponse)
	DeleteCustomerById(ctx context.Context, id uint64) (origin.DefaultResponse, int, origin.DefaultResponse)
}

type customerControllerStruct struct {
	customerRepository CustomerRepoInterface
}

func (c customerControllerStruct) CreateCustomer(ctx context.Context, req RequestCustomer) (origin.DefaultResponse, int, origin.DefaultResponse) {
	//var customerCustomer model.Customer
	var response origin.DefaultResponse
	var errorMessage origin.DefaultResponse

	// create acustomer
	status, err := c.customerRepository.CreateCustomer(ctx, req)
	if err != nil {
		errorMessage = origin.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	response = origin.DefaultSuccessResponseWithMessage("contoller_customer created", status, "success")
	return response, status, errorMessage
}

func (c customerControllerStruct) GetCustomerById(ctx context.Context, id uint64) (origin.DefaultResponse, int, origin.DefaultResponse) {
	var customerRepo model.Customer
	var response origin.DefaultResponse
	var errorMessage origin.DefaultResponse

	//get data by id
	status, err := c.customerRepository.GetCustomerById(ctx, id, &customerRepo)
	if err != nil {
		errorMessage = origin.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}
	response = origin.DefaultSuccessResponseWithMessage("Get contoller_customer", status, customerRepo)
	return response, status, errorMessage
}

func (c customerControllerStruct) GetAllCustomer(ctx context.Context, page uint64, req RequestGetAllCustomer) (origin.DefaultResponse, int, origin.DefaultResponse) {
	var customerRepo []model.Customer
	var customerCountRepo model.Customer
	var response origin.DefaultResponse
	var errorMessage origin.DefaultResponse

	var limit uint64 = 30
	status, err := c.customerRepository.GetCountRowsCustomer(ctx, &customerCountRepo)
	if err != nil {
		errorMessage = origin.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	status, err = c.customerRepository.GetAllCustomer(ctx, page, limit, req, &customerRepo)
	if err != nil {
		errorMessage = origin.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	resMessage := FindAllCustomer{
		Page:       page,
		PerPage:    uint64(len(customerRepo)),
		TotalPages: uint64(math.Ceil(float64(customerCountRepo.Total) / float64(limit))),
		Data:       customerRepo,
	}

	response = origin.DefaultSuccessResponseWithMessage("Get all contoller_customer", status, resMessage)

	return response, status, errorMessage
}

func (c customerControllerStruct) UpdateCustomerById(ctx context.Context, id uint64, req RequestUpdateCustomer) (origin.DefaultResponse, int, origin.DefaultResponse) {
	var customerRepo model.Customer
	var response origin.DefaultResponse
	var errorMessage origin.DefaultResponse

	//check authorization
	if id == 1 {
		errorMessage = origin.DefaultErrorResponseWithMessage("not authorization update", http.StatusUnauthorized)
		return response, http.StatusUnauthorized, errorMessage
	}

	//update data by id
	status, err := c.customerRepository.UpdateCustomerById(ctx, id, req)
	if err != nil {
		errorMessage = origin.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	//repo
	status, err = c.customerRepository.GetCustomerById(ctx, id, &customerRepo)
	if err != nil {
		errorMessage = origin.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}

	response = origin.DefaultSuccessResponseWithMessage("Get contoller_customer", status, customerRepo)
	return response, status, errorMessage
}

func (c customerControllerStruct) DeleteCustomerById(ctx context.Context, id uint64) (origin.DefaultResponse, int, origin.DefaultResponse) {
	var response origin.DefaultResponse
	var errorMessage origin.DefaultResponse
	//check authorization
	if id == 1 {
		errorMessage = origin.DefaultErrorResponseWithMessage("not authorization delete", http.StatusUnauthorized)
		return response, http.StatusUnauthorized, errorMessage
	}

	//repo
	status, err := c.customerRepository.DeleteCustomerById(ctx, id)
	if err != nil {
		errorMessage = origin.DefaultErrorResponseWithMessage(err.Error(), status)
		return response, status, errorMessage
	}
	response = origin.DefaultSuccessResponseWithMessage("delete contoller_customer", status, "true")
	return response, status, errorMessage
}
