package repository_customer

import (
	"context"
	"crm_service/app/clients/connection"
	"crm_service/app/config"
	"crm_service/app/model/model_customer"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type ClientRepositoryCustomer struct {
	client connection.InterfaceConnection
	conf   *config.Config
}

func NewClientCustomer(conf *config.Config, con connection.InterfaceConnection) InterfaceRepoCustomer {
	return &ClientRepositoryCustomer{
		client: con,
		conf:   conf,
	}
}
func (repo *ClientRepositoryCustomer) CreateCustomer(ctx context.Context, req model_customer.RequestCustomer) (int, error) {
	//timeout
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	var args []interface{}
	args = append(args, req.FirstName, req.LastName, req.Email, req.Avatar, time.Now().Format("20060102150405"), req.Email)
	//query
	queryCreateCustomer := "INSERT INTO customers( firstname, lastname,email,avatar,created_at) SELECT ?,?,?,?,? WHERE NOT EXISTS (SELECT email FROM customers WHERE email=?)"
	result := repo.client.GetConnectionDB().WithContext(ctx).Exec(queryCreateCustomer, args...)

	//check
	if result.Error != nil {
		return http.StatusInternalServerError, errors.New("failed exec query create queryCreateCustomer")
	} else if result.RowsAffected == 0 {
		// Username does not exist, proceed with creating the model
		return http.StatusInternalServerError, errors.New("email already exists")
	}

	//return
	return http.StatusCreated, nil
}

func (repo *ClientRepositoryCustomer) GetCustomerByEmail(ctx context.Context, req model_customer.RequestCustomerEmail) (status int, err error, res model_customer.Customer) {
	//timeout
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	var args []interface{}
	args = append(args, req.Email)
	//query
	queryCustomer := "SELECT id,firstname,lastname,email,avatar,created_at from customers WHERE email = ?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Raw(queryCustomer, args...).Scan(&res)

	//check
	if result.Error != nil {
		return http.StatusInternalServerError, errors.New("failed exec query create queryCustomer"), res
	} else if result.RowsAffected == 0 {
		// Username does not exist, proceed with creating the model
		return http.StatusInternalServerError, errors.New("customer not found"), res
	}

	//return
	return http.StatusCreated, nil, res
}

func (repo *ClientRepositoryCustomer) GetCustomerById(ctx context.Context, req uint64) (status int, err error, res model_customer.Customer) {
	//timeout
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	var args []interface{}
	args = append(args, req)
	//query
	queryCustomer := "SELECT id,firstname,lastname,email,avatar,created_at from customers WHERE id = ?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Raw(queryCustomer, args...).Scan(&res)

	//check
	if result.Error != nil {
		return http.StatusInternalServerError, errors.New("failed exec query create queryCustomer"), res
	} else if result.RowsAffected == 0 {
		// Username does not exist, proceed with creating the model
		return http.StatusInternalServerError, errors.New("customer not found"), res
	}

	//return
	return http.StatusOK, nil, res
}
func (repo *ClientRepositoryCustomer) GetCountRowsCustomer(ctx context.Context) (status int, err error, res model_customer.Customer) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	//query
	queryGetCustomerById := "SELECT count(id) AS total FROM customers"
	result := repo.client.GetConnectionDB().WithContext(ctx).Raw(queryGetCustomerById).Scan(&res)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query count all repository-model_customer"), res
	}

	return http.StatusOK, nil, res
}

func (repo *ClientRepositoryCustomer) GetAllCustomer(ctx context.Context, page uint64, limit uint64, firstname string, lastname string) (status int, err error, res []model_customer.Customer) {
	//timeout
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	//page
	startID := (page - 1) * limit
	var args []interface{}
	args = append(args, startID, fmt.Sprint("%", firstname, "%"), fmt.Sprint("%", lastname, "%"), limit)
	//query
	queryCustomer := "SELECT id,firstname,lastname,email,avatar,created_at from customers WHERE id > ? AND firstname LIKE ? OR lastname LIKE ? LIMIT ?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Raw(queryCustomer, args...).Scan(&res)

	//check
	if result.Error != nil {
		return http.StatusInternalServerError, errors.New("failed exec query create queryCustomer"), res
	} else if result.RowsAffected == 0 {
		// Username does not exist, proceed with creating the model
		return http.StatusInternalServerError, errors.New("customer not found"), res
	}

	//return
	return http.StatusOK, nil, res
}

func (repo *ClientRepositoryCustomer) UpdateCustomerById(ctx context.Context, id uint64, req model_customer.RequestUpdateCustomer) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	var args []interface{}
	args = append(args, req.FirstName, req.LastName, req.Avatar, time.Now().Format("20060102150405"), id)
	//query
	queryUpdateCustomerById := "UPDATE customers SET firstname=?,lastname=?,avatar=?,updated_at=? WHERE id=?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Exec(queryUpdateCustomerById, args...)
	if result.Error != nil {
		//username already exist
		if strings.Contains(result.Error.Error(), "Error 1062 (23000)") {
			return http.StatusBadRequest, errors.New("email already exist")
		}
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query UpdateCustomerById")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("data not found")
	}
	return http.StatusAccepted, nil
}

func (repo *ClientRepositoryCustomer) DeleteCustomerById(ctx context.Context, id uint64) (int, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(repo.conf.Database.Timeout)*time.Millisecond)
	defer cancel()

	var args []interface{}
	args = append(args, id)

	queryDeleteCustomerById := "DELETE FROM customers WHERE id =?"
	result := repo.client.GetConnectionDB().WithContext(ctx).Exec(queryDeleteCustomerById, args...)
	if result.Error != nil {
		//error mysql
		return http.StatusInternalServerError, errors.New("failed exec query DeleteCustomerById")
	} else if result.RowsAffected == 0 {
		// return if not found
		return http.StatusNotFound, errors.New("customer is not found,delete unacceptable")
	}
	return http.StatusOK, nil
}
