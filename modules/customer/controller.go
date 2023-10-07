package customer

import (
	"crm_service/entity"

	"crm_service/model"
	"fmt"
	"time"
)

type CustomerControllerInterface interface {
	CreateCustomer(req CustomerBody) (any, error)
	GetCustomerById(id uint) (FindCustomer, error)
	GetAllCustomer(page uint, usernameStr string) (FindAllCustomer, error)
	UpdateById(id uint, req UpdateCustomerBody) (FindCustomer, error)
	DeleteCustomerById(id uint) (entity.ResponseMeta, error)
}

type customerControllerStruct struct {
	customerUseCase UseCaseCustomerInterface
}

func (c customerControllerStruct) CreateCustomer(req CustomerBody) (any, error) {
	start := time.Now()
	customer, err := c.customerUseCase.CreateCustomer(req)
	if err != nil {
		return SuccessCreate{}, err
	}

	res := SuccessCreate{
		ResponseMeta: entity.ResponseMeta{
			Success:      true,
			MessageTitle: "Success create customer",
			Message:      "Success register",
			ResponseTime: fmt.Sprint(time.Since(start)),
		},
		Data: CustomerBody{
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Email,
			Avatar:    customer.Avatar,
		},
	}
	return res, nil
}

func (c customerControllerStruct) GetCustomerById(id uint) (FindCustomer, error) {
	start := time.Now()
	var res FindCustomer
	customer, err := c.customerUseCase.GetCustomerById(id)
	if err != nil {
		return FindCustomer{}, err
	}

	res.ResponseMeta = entity.ResponseMeta{
		Success:      true,
		MessageTitle: "Success find customer",
		Message:      "Success find",
		ResponseTime: fmt.Sprint(time.Since(start)),
	}
	res.Data = customer
	return res, nil
}

func (c customerControllerStruct) GetAllCustomer(page uint, usernameStr string) (FindAllCustomer, error) {
	start := time.Now()
	page, perPage, total, totalPages, customerEntities, err := c.customerUseCase.GetAllCustomer(page, usernameStr)

	if err != nil {
		return FindAllCustomer{}, err
	}

	data := make([]model.Customer, len(customerEntities))
	for i, customerEntity := range customerEntities {
		data[i] = customerEntity
	}

	res := FindAllCustomer{
		ResponseMeta: entity.ResponseMeta{
			Success:      true,
			MessageTitle: "Success find customer",
			Message:      "Success find all",
			ResponseTime: fmt.Sprint(time.Since(start)),
		},
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: totalPages,
		Data:       data,
	}

	return res, nil
}

func (c customerControllerStruct) UpdateById(id uint, req UpdateCustomerBody) (FindCustomer, error) {
	start := time.Now()
	customer, err := c.customerUseCase.UpdateCustomerById(id, req)
	if err != nil {
		return FindCustomer{}, err
	}

	res := FindCustomer{
		ResponseMeta: entity.ResponseMeta{
			Success:      true,
			MessageTitle: "Success update customer",
			Message:      "Success update customer",
			ResponseTime: fmt.Sprint(time.Since(start)),
		},
		Data: customer,
	}
	return res, nil
}

func (c customerControllerStruct) DeleteCustomerById(id uint) (entity.ResponseMeta, error) {
	start := time.Now()
	err := c.customerUseCase.DeleteCustomerById(id)
	res := entity.ResponseMeta{
		Success:      true,
		MessageTitle: "Success delete customer",
		Message:      "Success delete customer",
		ResponseTime: fmt.Sprint(time.Since(start)),
	}
	return res, err
}
