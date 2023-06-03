package repository

import (
	"crm_service/entity"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math"
)

type CustomerRepoInterface interface {
	CreateCustomer(customer *entity.Customer) (entity.Customer, error)
	GetCustomerById(id uint) (entity.Customer, error)
	GetAllCustomer(page uint, username string) (uint, uint, int, uint, []entity.Customer, error)
	UpdateCustomerById(id uint, customer *entity.Customer) (entity.Customer, error)
	DeleteCustomerById(id uint) error
}

type Customer struct {
	db *gorm.DB
}

func NewCustomer(dbCrud *gorm.DB) Customer {
	return Customer{
		db: dbCrud,
	}

}

func (repo Customer) CreateCustomer(customer *entity.Customer) (entity.Customer, error) {
	var existingCustomer entity.Customer

	err := repo.db.First(&existingCustomer, "username = ?", customer.FirstName).Error
	if err == nil {
		// FirstName already exists, return an error
		return entity.Customer{}, errors.New("username already taken")
	}

	// FirstName does not exist, proceed with creating the customer
	err = repo.db.Create(customer).Error
	if err != nil {
		return entity.Customer{}, err
	}
	return *customer, nil
}

func (repo Customer) GetCustomerById(id uint) (entity.Customer, error) {
	var customer entity.Customer
	err := repo.db.Omit("password").First(&customer, "id = ?", id).Error
	if err != nil {
		return entity.Customer{}, errors.New("customer not found")
	}
	return customer, nil
}

func (repo Customer) GetAllCustomer(page uint, username string) (uint, uint, int, uint, []entity.Customer, error) {
	var customers []entity.Customer
	var count int64
	var limit uint = 20
	var offset = limit * (page - 1)
	result := repo.db.Model(&entity.Customer{}).Count(&count)
	if result.Error != nil {
		// Handle the error
		return 0, 0, 0, 0, nil, result.Error
	}
	totalPages := uint(math.Ceil(float64(count) / float64(limit)))
	err := repo.db.Omit("password").Limit(int(limit)).Offset(int(offset)).Where("username LIKE ?", fmt.Sprint("%", username, "%")).Find(&customers).Error
	if err != nil {
		return 0, 0, 0, 0, nil, err
	}
	return page, limit, int(count), totalPages, customers, nil
}

func (repo Customer) UpdateCustomerById(id uint, updateCustomer *entity.Customer) (entity.Customer, error) {
	var findCustomerById entity.Customer
	var existingCustomer entity.Customer

	if id == 1 {
		return entity.Customer{}, errors.New("customer is super admin and cannot be updated")
	}

	err := repo.db.First(&findCustomerById, "id = ?", id).Error
	if err != nil {
		return entity.Customer{}, errors.New("customer not found")
	}

	err = repo.db.Where("username = ?", updateCustomer.FirstName).Not("username = ?", findCustomerById.FirstName).First(&existingCustomer).Error
	fmt.Println(existingCustomer)
	if err == nil {
		// FirstName already exists, return an error
		return entity.Customer{}, errors.New("username already taken")
	}

	err = repo.db.Model(&entity.Customer{}).Where("id = ?", id).Updates(updateCustomer).Error
	if err != nil {
		return entity.Customer{}, errors.New("failed to update customer")
	}

	err = repo.db.First(&findCustomerById, "id = ?", id).Error
	if err != nil {
		return entity.Customer{}, errors.New("customer not found")
	}

	return findCustomerById, nil
}

func (repo Customer) DeleteCustomerById(id uint) error {
	var customer entity.Customer
	if id == 1 {
		return errors.New("customer is super admin cannot delete")
	}

	err := repo.db.First(&customer, "id = ?", id).Error
	if err != nil {
		return errors.New("customer not found")
	}
	err = repo.db.Delete(&customer, "id = ?", id).Error
	if err != nil {
		return errors.New("failed deleted")
	}
	return nil
}
