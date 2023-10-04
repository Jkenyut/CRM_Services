package customer

import (
	"crm_service/model"
	"crm_service/repository"
)

type UseCaseCustomerInterface interface {
	CreateCustomer(customer CustomerBody) (model.Customer, error)
	GetCustomerById(id uint) (model.Customer, error)
	GetAllCustomer(page uint, username string) (uint, uint, int, uint, []model.Customer, error)
	UpdateCustomerById(id uint, customer UpdateCustomerBody) (model.Customer, error)
	DeleteCustomerById(id uint) error
}

type customerUseCaseStruct struct {
	customerRepository repository.CustomerRepoInterface
}

func (uc customerUseCaseStruct) CreateCustomer(customer CustomerBody) (model.Customer, error) {

	NewCustomer := model.Customer{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Avatar:    customer.Avatar,
	}

	createCustomer, err := uc.customerRepository.CreateCustomer(&NewCustomer)
	if err != nil {
		return NewCustomer, err
	}
	return createCustomer, nil
}

func (uc customerUseCaseStruct) GetCustomerById(id uint) (model.Customer, error) {
	var customer model.Customer
	customer, err := uc.customerRepository.GetCustomerById(id)
	return customer, err
}

func (uc customerUseCaseStruct) GetAllCustomer(page uint, username string) (uint, uint, int, uint, []model.Customer, error) {
	var customer []model.Customer
	page, perPage, total, totalPages, customer, err := uc.customerRepository.GetAllCustomer(page, username)
	return page, perPage, total, totalPages, customer, err
}

func (uc customerUseCaseStruct) UpdateCustomerById(id uint, customer UpdateCustomerBody) (model.Customer, error) {

	newCustomer := model.Customer{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Avatar:    customer.Avatar,
	}

	updatedCustomer, err := uc.customerRepository.UpdateCustomerById(id, &newCustomer)
	if err != nil {
		return newCustomer, err
	}

	return updatedCustomer, nil
}

func (uc customerUseCaseStruct) DeleteCustomerById(id uint) error {
	err := uc.customerRepository.DeleteCustomerById(id)
	return err
}
