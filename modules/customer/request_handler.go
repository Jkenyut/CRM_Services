package customer

import (
	"crm_service/entity"
	"crm_service/utils/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RequestHandlerCustomerStruct struct {
	ctr CustomerControllerInterface
}

func RequestHandler(
	dbCrud *gorm.DB,
) RequestHandlerCustomerStruct {
	return RequestHandlerCustomerStruct{
		ctr: customerControllerStruct{
			customerRepository: NewCustomer(dbCrud),
		}}
}

var validate = validator.New()

func (h RequestHandlerCustomerStruct) CreateCustomer(c *gin.Context) {
	var request RequestCustomer
	err := c.Bind(&request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, entity.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
		return
	}

	//validate
	err = validate.Struct(request)
	if err != nil {
		// Validation failed
		c.AbortWithStatusJSON(helper.ValidateData(err))
		return

	}
	//controller
	res, status, errMessage := h.ctr.CreateCustomer(c, request)
	//check status
	if status < 200 || status > 299 {
		c.AbortWithStatusJSON(status, errMessage)
		return
	}
	c.JSON(http.StatusCreated, res)
}
func (h RequestHandlerCustomerStruct) GetCustomerById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, entity.DefaultErrorResponseWithMessage("must unsigned number", http.StatusBadRequest))
		return
	}
	res, status, errMessage := h.ctr.GetCustomerById(c, id)
	//check status
	if status < 200 || status > 299 {
		c.AbortWithStatusJSON(status, errMessage)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerCustomerStruct) GetAllCustomer(c *gin.Context) {
	page, err := strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 64)
	fisrtName := c.DefaultQuery("firstname", "")
	lastName := c.DefaultQuery("lastname", "")
	requestBody := RequestGetAllCustomer{
		FirstName: fisrtName,
		LastName:  lastName,
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, entity.DefaultErrorResponseWithMessage("must unsigned number", http.StatusBadRequest))
		return
	}

	res, status, errMessage := h.ctr.GetAllCustomer(c, page, requestBody)
	//check status
	if status < 200 || status > 299 {
		c.AbortWithStatusJSON(status, errMessage)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerCustomerStruct) UpdateCustomerById(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, entity.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
		return
	}

	request := RequestUpdateCustomer{}
	err = c.Bind(&request)
	//validate
	err = validate.Struct(request)
	if err != nil {
		// Validation failed
		c.AbortWithStatusJSON(helper.ValidateData(err))
		return
	}

	res, status, errMessage := h.ctr.UpdateCustomerById(c, id, request)
	//check status
	if status < 200 || status > 299 {
		c.AbortWithStatusJSON(status, errMessage)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerCustomerStruct) DeleteCustomerById(c *gin.Context) {

	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, entity.DefaultErrorResponseWithMessage("required not valid", http.StatusBadRequest))
		return
	}

	res, status, errMessage := h.ctr.DeleteCustomerById(c, actorId)
	//check status
	if status < 200 || status > 299 {
		c.AbortWithStatusJSON(status, errMessage)
		return
	}
	c.JSON(http.StatusOK, res)
}
