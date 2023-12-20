package controller_customer

import (
	"crm_service/app/clients/repository/repository_customer"
	"crm_service/app/middleware/pipeline"
	"crm_service/app/model/model_customer"
	"crm_service/app/utils/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

type ControllerCustomer struct {
	client    repository_customer.InterfaceRepoCustomer
	validator *validator.Validate
}

func NewControllerCustomer(client repository_customer.InterfaceRepoCustomer, validate *validator.Validate) InterfaceControllerCustomer {
	return &ControllerCustomer{
		client:    client,
		validator: validate,
	}
}

func (ctr *ControllerCustomer) CreateCustomer(c *gin.Context) {
	var request model_customer.RequestCustomer
	if valid := pipeline.BindAndValidateRequest(c, ctr.validator, &request); valid {
		return // Error response handled in bindAndValidateRequest
	}

	status, err := ctr.client.CreateCustomer(c, request)
	if err != nil || !helper.IsSuccessStatus(status) {
		pipeline.AbortWithStatusJSON(c, status, err.Error())
		return
	}

	response := model_customer.ResponseCustomer{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Avatar:    request.Avatar,
	}

	pipeline.JSON(c, status, "customer created", response)
}

func (ctr *ControllerCustomer) GetCustomerByEmail(c *gin.Context) {
	var request model_customer.RequestCustomerEmail
	if valid := pipeline.BindAndValidateRequest(c, ctr.validator, &request); valid {
		return // Error response handled in bindAndValidateRequest
	}

	status, err, res := ctr.client.GetCustomerByEmail(c, request)
	if err != nil || !helper.IsSuccessStatus(status) {
		pipeline.AbortWithStatusJSON(c, status, err.Error())
		return
	}

	response := model_customer.ResponseCustomer{
		ID:        res.ID,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Email:     res.Email,
		Avatar:    res.Avatar,
		CreatedAt: helper.ConvertTimeToWIB(res.CreatedAt),
	}

	pipeline.JSON(c, status, "get customer by email", response)
}

func (ctr *ControllerCustomer) GetCustomerById(c *gin.Context) {
	//bind param
	id, valid := pipeline.BindParamAndParseUint(c, "id")
	if valid {
		return
	}

	status, err, res := ctr.client.GetCustomerById(c, id)
	if err != nil || !helper.IsSuccessStatus(status) {
		pipeline.AbortWithStatusJSON(c, status, err.Error())
		return
	}

	response := model_customer.ResponseCustomer{
		ID:        res.ID,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Email:     res.Email,
		Avatar:    res.Avatar,
		CreatedAt: helper.ConvertTimeToWIB(res.CreatedAt),
	}

	pipeline.JSON(c, status, "get customer by id", response)
}

func (ctr *ControllerCustomer) GetAllCustomer(c *gin.Context) {

	page, valid := pipeline.BindQueryAndParseUint(c, "page", "1")
	limit, valid := pipeline.BindQueryAndParseUint(c, "limit", "10")
	if valid {
		return // if error
	}

	firstname := c.DefaultQuery("firstname", "")
	lastname := c.DefaultQuery("lastname", "")
	status, err, allActor := ctr.client.GetAllCustomer(c, page, limit, firstname, lastname)
	//check status
	if err != nil || !helper.IsSuccessStatus(status) {
		pipeline.AbortWithStatusJSON(c, status, err.Error())
		return
	}
	status, err, count := ctr.client.GetCountRowsCustomer(c)
	//check status
	if err != nil || !helper.IsSuccessStatus(status) {
		pipeline.AbortWithStatusJSON(c, status, err.Error())
		return
	}
	var customer []model_customer.ResponseCustomer
	for _, i := range allActor {
		history := model_customer.ResponseCustomer{
			ID:        i.ID,
			FirstName: i.FirstName,
			LastName:  i.LastName,
			Email:     i.Email,
			Avatar:    i.Avatar,
			CreatedAt: helper.ConvertTimeToWIB(i.CreatedAt),
		}
		customer = append(customer, history)
	}

	response := model_customer.FindAllCustomer{
		Page:       page,
		PerPage:    limit,
		TotalPages: helper.CustomFloor(float64(count.Total / limit)),
		TotalData:  count.Total,
		Data:       customer,
	}
	pipeline.JSON(c, http.StatusOK, "Get All Actor", response)
}

func (ctr *ControllerCustomer) UpdateCustomerById(c *gin.Context) {

	//bind param
	id, valid := pipeline.BindParamAndParseUint(c, "id")
	if valid {
		return
	}

	var request model_customer.RequestUpdateCustomer
	if valid = pipeline.BindAndValidateRequest(c, ctr.validator, &request); valid {
		return // Error response handled in bindAndValidateRequest
	}

	status, err := ctr.client.UpdateCustomerById(c, id, request)
	//check status
	if err != nil || !helper.IsSuccessStatus(status) {
		pipeline.AbortWithStatusJSON(c, status, err.Error())
		return
	}
	response := model_customer.ResponseCustomer{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Avatar:    request.Avatar,
		UpdatedAt: helper.ConvertTimeToWIB(time.Now()),
	}
	pipeline.JSON(c, status, "update actor", response)
}
func (ctr *ControllerCustomer) DeleteCustomerById(c *gin.Context) {

	//bind param
	id, valid := pipeline.BindParamAndParseUint(c, "id")
	if valid {
		return
	}

	status, err := ctr.client.DeleteCustomerById(c, id)
	//check status
	if err != nil || !helper.IsSuccessStatus(status) {
		pipeline.AbortWithStatusJSON(c, status, err.Error())
		return
	}
	pipeline.JSON(c, status, "Delete customer", "Success")
}
