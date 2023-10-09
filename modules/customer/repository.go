package customer

import (
	"context"
	"crm_service/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type CustomerRepoInterface interface {
	CreateCustomer(ctx context.Context, req RequestCustomer) (int, error)
	GetCustomerByEmail(ctx context.Context, req RequestCustomer, customerRepository *model.Customer) (int, error)
	GetCustomerById(ctx context.Context, id uint64, customerRepository *model.Customer) (int, error)
	GetAllCustomer(ctx context.Context, page uint64, limit uint64, req RequestGetAllCustomer, customerRepository *[]model.Customer) (int, error)
	GetCountRowsCustomer(ctx context.Context, customerRepository *model.Customer) (int, error)
	UpdateCustomerById(ctx context.Context, id uint64, customerRepository RequestUpdateCustomer) (int, error)
	DeleteCustomerById(ctx context.Context, id uint64) (int, error)
}

type Customer struct {
	db *gorm.DB
}

func NewCustomer(dbCrud *gorm.DB) Customer {
	return Customer{
		db: dbCrud,
	}

}

func (repo Customer) CreateCustomer(ctx context.Context, req RequestCustomer) (int, error) {
	//timeout
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(3000)*time.Millisecond)
	defer cancel()
	//query
	queryCreateCustomer := "INSERT INTO customer(first_name, last_name, email, avatar)  SELECT ?,?,?,? WHERE NOT EXISTS (SELECT email from customer where email=?)"
	result := repo.db.WithContext(ctx).Exec(queryCreateCustomer, req.FirstName, req.LastName, req.Email, req.Avatar, req.Email)

	//check
	if result.Error != nil {
		return http.StatusInternalServerError, errors.New("failed exec query create customer")
	} else if result.RowsAffected == 0 {
		// Username does not exist, proceed with creating the customer
		return http.StatusInternalServerError, errors.New("username already exists")
	}

	//return
	return http.StatusCreated, nil
}

func (repo Customer) GetCustomerByEmail(ctx context.Context, req RequestCustomer, customerRepository *model.Customer) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(3000)*time.Millisecond)
	defer cancel()

	//query
	querySelectCustomer := "select id, first_name, last_name, email, avatar, created_at, updated_at from customer where email=?"
	result := repo.db.WithContext(ctx).Raw(querySelectCustomer, req.Email).Scan(&customerRepository)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query login customer")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("customer not found")
	}

	return http.StatusOK, nil
}

func (repo Customer) GetCustomerById(ctx context.Context, id uint64, customerRepository *model.Customer) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(3000)*time.Millisecond)
	defer cancel()

	//query
	queryGetCustomerById := "select id, first_name, last_name, email, avatar, created_at, updated_at from customer where id=?"
	result := repo.db.WithContext(ctx).Raw(queryGetCustomerById, id).Scan(&customerRepository)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query login customer")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("customer not found")
	}

	return http.StatusOK, nil
}

func (repo Customer) GetAllCustomer(ctx context.Context, page uint64, limit uint64, req RequestGetAllCustomer, customerRepository *[]model.Customer) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(3000)*time.Millisecond)
	defer cancel()

	//page
	startID := (page - 1) * limit

	//query
	queryGetCustomerById := "select id, first_name, last_name, email, avatar, created_at, updated_at from customer where id > ? AND first_name like ? AND last_name like ? limit ?"
	result := repo.db.WithContext(ctx).Raw(queryGetCustomerById, startID, fmt.Sprint(req.FirstName, "%"), fmt.Sprint(req.LastName, "%"), limit).Scan(&customerRepository)

	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query all customer")
	}

	return http.StatusOK, nil
}
func (repo Customer) GetCountRowsCustomer(ctx context.Context, customerRepository *model.Customer) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(3000)*time.Millisecond)
	defer cancel()

	//query
	queryGetCustomerById := "select count(id) as total from customer"
	result := repo.db.WithContext(ctx).Raw(queryGetCustomerById).Scan(&customerRepository)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query login customer")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("count customer not found")
	}

	return http.StatusOK, nil
}

func (repo Customer) UpdateCustomerById(ctx context.Context, id uint64, updateCustomer RequestUpdateCustomer) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(3000)*time.Millisecond)
	defer cancel()

	//query
	queryUpdateCustomerById := "update customer set  first_name=?,last_name=?,avatar=? WHERE id=?"
	result := repo.db.WithContext(ctx).Exec(queryUpdateCustomerById, updateCustomer.FirstName, updateCustomer.LastName, updateCustomer.Avatar, id)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query UpdateCustomerById")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("cannot update because username already exist")
	}

	return http.StatusOK, nil
}

func (repo Customer) DeleteCustomerById(ctx context.Context, id uint64) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(3000)*time.Millisecond)
	defer cancel()

	queryDeleteCustomerById := "delete from customer where id =?"
	result := repo.db.WithContext(ctx).Exec(queryDeleteCustomerById, id)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query DeleteCustomerById")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("customer is not found,delete unacceptable")
	}
	return http.StatusOK, nil
}
