package customer

import (
	"crm_service/dto"
	"crm_service/repository"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestHandlerCustomerStruct struct {
	ctr CustomerControllerInterface
}

func RequestHandler(
	dbCrud *gorm.DB,
) RequestHandlerCustomerStruct {
	return RequestHandlerCustomerStruct{
		ctr: customerControllerStruct{
			customerUseCase: customerUseCaseStruct{
				customerRepository: repository.NewCustomer(dbCrud),
			},
		}}
}

var validate = validator.New()

func (h RequestHandlerCustomerStruct) CreateCustomer(c *gin.Context) {
	request := CustomerBody{}
	err := c.Bind(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	err = validate.Struct(request)

	if err != nil {
		// Validation failed

		for _, err := range err.(validator.ValidationErrors) {
			customErr := fmt.Sprint(err.StructField(), " ", err.ActualTag(), " ", err.Param())
			switch err.Tag() {
			case "required":
				c.JSON(http.StatusUnprocessableEntity, dto.DefaultErrorResponseWithMessage(customErr))
				return
			case "min":
				c.JSON(http.StatusUnprocessableEntity, dto.DefaultErrorResponseWithMessage(customErr))
				return
			case "max":
				c.JSON(http.StatusUnprocessableEntity, dto.DefaultErrorResponseWithMessage(customErr))
				return
			case "email":
				c.JSON(http.StatusUnprocessableEntity, dto.DefaultErrorResponseWithMessage("email wrong"))
				return
			case "alphanum":
				c.JSON(http.StatusUnprocessableEntity, dto.DefaultErrorResponseWithMessage(customErr))
				return
			case "alphanumunicode":
				c.JSON(http.StatusUnprocessableEntity, dto.DefaultErrorResponseWithMessage(customErr))
				return
			}
		}
	}
	res, err := h.ctr.CreateCustomer(request)
	if err != nil {
		if err.Error() == "email already taken" {
			c.JSON(http.StatusConflict, dto.DefaultErrorResponseWithMessage("email already taken"))
			return
		} else {
			c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage("Server error"))
			return
		}
	}
	c.JSON(http.StatusCreated, res)
}

func (h RequestHandlerCustomerStruct) GetCustomerById(c *gin.Context) {
	customerId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.GetCustomerById(uint(customerId))
	if err != nil {
		if err.Error() == "customer not found" {
			c.JSON(http.StatusNotFound, dto.DefaultErrorResponseWithMessage("Customer not found"))
			return
		} else {
			c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage("Server error"))
			return
		}

	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerCustomerStruct) GetAllCustomer(c *gin.Context) {

	pageStr := c.DefaultQuery("page", "1")
	usernameStr := c.DefaultQuery("name", "")
	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.GetAllCustomer(uint(page), usernameStr)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerCustomerStruct) UpdateCustomerById(c *gin.Context) {
	request := UpdateCustomerBody{}
	err := c.Bind(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	customerId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	err = validate.Struct(request)

	if err != nil {
		// Validation failed

		for _, err := range err.(validator.ValidationErrors) {
			customErr := fmt.Sprint(err.StructField(), " ", err.ActualTag(), " ", err.Param())
			switch err.Tag() {
			case "required":
				c.JSON(http.StatusUnprocessableEntity, dto.DefaultErrorResponseWithMessage(customErr))
				return
			case "min":
				c.JSON(http.StatusUnprocessableEntity, dto.DefaultErrorResponseWithMessage(customErr))
				return
			case "max":
				c.JSON(http.StatusUnprocessableEntity, dto.DefaultErrorResponseWithMessage(customErr))
				return
			case "alphanum":
				c.JSON(http.StatusUnprocessableEntity, dto.DefaultErrorResponseWithMessage(customErr))
				return
			case "eq":
				c.JSON(http.StatusUnprocessableEntity, dto.DefaultErrorResponseWithMessage(customErr))
				return

			}
		}
	}
	res, err := h.ctr.UpdateById(uint(customerId), request)
	if err != nil {
		if err.Error() == "customer not found" {
			c.JSON(http.StatusNotFound, dto.DefaultErrorResponseWithMessage("customer not found"))
			return
		} else if err.Error() == "customer is super admin cannot update" {
			c.JSON(http.StatusUnauthorized, dto.DefaultErrorResponseWithMessage("customer is super admin cannot update"))
			return
		} else if err.Error() == "username already taken" {
			c.JSON(http.StatusConflict, dto.DefaultErrorResponseWithMessage("username already taken"))
			return
		} else if err.Error() == "failed to update customer" {
			c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage("failed to update customer"))
			return
		} else {
			c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage("Server error"))
			return
		}
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerCustomerStruct) DeleteCustomerById(c *gin.Context) {
	customerId, err := strconv.ParseUint(c.Param("id"), 10, 64)

	res, err := h.ctr.DeleteCustomerById(uint(customerId))
	if err != nil {
		if err.Error() == "customer not found" {
			c.JSON(http.StatusNotFound, dto.DefaultErrorResponseWithMessage("Customer not found"))
			return
		} else if err.Error() == "customer is super admin cannot delete" {
			c.JSON(http.StatusUnauthorized, dto.DefaultErrorResponseWithMessage("customer is super admin cannot delete"))
			return
		} else if err.Error() == "failed deleted" {
			c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage("failed deleted"))
			return
		} else {
			c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage("Server error"))
			return
		}

	}
	c.JSON(http.StatusOK, res)
}
