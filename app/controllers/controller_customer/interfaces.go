package controller_customer

import "github.com/gin-gonic/gin"

type InterfaceControllerCustomer interface {
	CreateCustomer(c *gin.Context)
	GetCustomerByEmail(c *gin.Context)
	GetCustomerById(c *gin.Context)
	GetAllCustomer(c *gin.Context)
	UpdateCustomerById(c *gin.Context)
	DeleteCustomerById(c *gin.Context)
}
