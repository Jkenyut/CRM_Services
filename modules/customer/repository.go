package customer

import (
	"crm_service/model"
	"errors"
	"gorm.io/gorm"
)

type CustomerRepoInterface interface {
	CreateCustomer(customer *model.Customer) (model.Customer, error)
	//GetCustomerById(id uint) (model.Customer, error)
	//GetAllCustomer(page uint, username string) (uint, uint, int, uint, []model.Customer, error)
	//UpdateCustomerById(id uint, customer *model.Customer) (model.Customer, error)
	//DeleteCustomerById(id uint) error
}

type Customer struct {
	db *gorm.DB
}

func NewCustomer(dbCrud *gorm.DB) Customer {
	return Customer{
		db: dbCrud,
	}

}

func (repo Customer) CreateCustomer(customer *model.Customer) (model.Customer, error) {
	var existingCustomer model.Customer

	err := repo.db.First(&existingCustomer, "email = ?", customer.Email).Error
	if err == nil {
		// FirstName already exists, return an error
		return model.Customer{}, errors.New("email already taken")
	}

	// FirstName does not exist, proceed with creating the customer
	err = repo.db.Create(customer).Error
	if err != nil {
		return model.Customer{}, err
	}
	return *customer, nil
}

func (repo Customer) GetCustomerById(id uint) (model.Customer, error) {
	var customer model.Customer
	err := repo.db.Omit("password").First(&customer, "id = ?", id).Error
	if err != nil {
		return model.Customer{}, errors.New("customer not found")
	}
	return customer, nil
}

//
//func (repo Customer) GetAllCustomer(page uint, username string) (uint, uint, int, uint, []model.Customer, error) {
//	var customers []model.Customer
//	var count int64
//	var limit uint = 20
//	var offset = limit * (page - 1)
//	result := repo.db.Model(&model.Customer{}).Count(&count)
//	if result.Error != nil {
//		// Handle the error
//		return 0, 0, 0, 0, nil, result.Error
//	}
//	totalPages := uint(math.Ceil(float64(count) / float64(limit)))
//	name := fmt.Sprintf("%%%s%%", username)
//
//	err := repo.db.Select("*").
//		Table("customer").
//		Select("*").
//		Where("CONCAT(first_name, ' ', last_name) LIKE ?", name).
//		Limit(int(limit)).
//		Offset(int(offset)).
//		Find(&customers).
//		Error
//	if err != nil {
//		return 0, 0, 0, 0, nil, err
//	}
//	return page, limit, int(count), totalPages, customers, nil
//}
//
//func (repo Customer) UpdateCustomerById(id uint, updateCustomer *model.Customer) (model.Customer, error) {
//	var findCustomerById model.Customer
//
//	err := repo.db.First(&findCustomerById, "id = ?", id).Error
//	if err != nil {
//		return model.Customer{}, errors.New("customer not found")
//	}
//
//	err = repo.db.Model(&model.Customer{}).Where("id = ?", id).Updates(updateCustomer).Error
//	if err != nil {
//		return model.Customer{}, errors.New("failed to update customer")
//	}
//
//	err = repo.db.First(&findCustomerById, "id = ?", id).Error
//	if err != nil {
//		return model.Customer{}, errors.New("customer not found")
//	}
//
//	return findCustomerById, nil
//}
//
//func (repo Customer) DeleteCustomerById(id uint) error {
//	var customer model.Customer
//
//	err := repo.db.First(&customer, "id = ?", id).Error
//	if err != nil {
//		return errors.New("customer not found")
//	}
//	err = repo.db.Delete(&customer, "id = ?", id).Error
//	if err != nil {
//		return errors.New("failed deleted")
//	}
//	return nil
//}
