package contoller_actor

import (
	"crm_service/app/clients/repository/repository_actor"
	"crm_service/app/middleware/pipeline"
	"crm_service/app/model/model_actor"
	"crm_service/app/utils/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type ControllerActor struct {
	client    repository_actor.InterfaceRepositoryActor
	validator *validator.Validate
}

func NewControllerActor(client repository_actor.InterfaceRepositoryActor, validate *validator.Validate) InterfaceControllerActor {
	return &ControllerActor{
		client:    client,
		validator: validate,
	}
}

const bcryptDefaultCost = 12
const defaultRoleID = 2

func (ctr *ControllerActor) CreateActor(c *gin.Context) {
	if valid := pipeline.ValidateJWT(c); valid {
		return // Error response handled in validateJWT
	}

	var request model_actor.RequestActor
	if valid := pipeline.BindAndValidateRequest(c, ctr.validator, &request); valid {
		return // Error response handled in bindAndValidateRequest
	}

	hashingPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcryptDefaultCost)
	if err != nil {
		pipeline.AbortWithStatusJSON(c, http.StatusPreconditionFailed, "Precondition failed Request")
		return
	}

	request.Password = string(hashingPassword)

	status, err := ctr.client.CreateActor(c, &request)
	if err != nil || !helper.IsSuccessStatus(status) {
		pipeline.AbortWithStatusJSON(c, status, err.Error())
		return
	}

	response := model_actor.ResponseActor{
		Username:  request.Username,
		CreatedAt: helper.ConvertTimeToWIB(time.Now()),
		RoleID:    defaultRoleID,
		Active:    "false",
	}

	pipeline.JSON(c, status, "actor created", response)
}

func (ctr *ControllerActor) GetActorById(c *gin.Context) {
	if valid := pipeline.ValidateJWT(c); valid {
		return // Error response handled in validateJWT
	}

	//bind param
	id, valid := pipeline.BindParamAndParseUint(c, "id")
	if valid {
		return
	}

	status, err, getActor := ctr.client.GetActorById(c, id)
	//check status
	if err != nil || !helper.IsSuccessStatus(status) {
		pipeline.AbortWithStatusJSON(c, status, err.Error())
		return
	}
	response := model_actor.ResponseActor{
		ID:        getActor.ID,
		Username:  getActor.Username,
		RoleID:    getActor.RoleID,
		Active:    getActor.Active,
		Verified:  getActor.Verified,
		CreatedAt: helper.ConvertTimeToWIB(getActor.CreatedAt),
		UpdatedAt: helper.ConvertTimeToWIB(getActor.UpdatedAt),
	}

	pipeline.JSON(c, status, "Get Actor", response) // res
}

func (ctr *ControllerActor) GetAllActor(c *gin.Context) {

	page, valid := pipeline.BindQueryAndParseUint(c, "page")
	limit, valid := pipeline.BindQueryAndParseUint(c, "limit")
	if valid {
		return // if error
	}

	username := c.DefaultQuery("username", "")
	status, err, allActor := ctr.client.GetAllActor(c, page, limit, username)
	//check status
	if err != nil || !helper.IsSuccessStatus(status) {
		pipeline.AbortWithStatusJSON(c, status, err.Error())
		return
	}
	status, err, countActor := ctr.client.GetCountRowsActor(c)
	//check status
	if err != nil || !helper.IsSuccessStatus(status) {
		pipeline.AbortWithStatusJSON(c, status, err.Error())
		return
	}
	var actor []model_actor.ResponseActor
	for _, i := range allActor {
		history := model_actor.ResponseActor{
			ID:        i.ID,
			Username:  i.Username,
			RoleID:    i.RoleID,
			Active:    i.Active,
			Verified:  i.Verified,
			CreatedAt: helper.ConvertTimeToWIB(i.CreatedAt),
			UpdatedAt: helper.ConvertTimeToWIB(i.UpdatedAt),
		}
		actor = append(actor, history)
	}

	response := model_actor.FindAllActor{
		Page:       page,
		PerPage:    limit,
		TotalPages: countActor.Total,
		Data:       actor,
	}
	pipeline.JSON(c, http.StatusOK, "Get All Actor", response)
}

func (ctr *ControllerActor) UpdateActorById(c *gin.Context) {
	if valid := pipeline.ValidateJWT(c); valid {
		return // Error response handled in validateJWT
	}

	//bind param
	id, valid := pipeline.BindParamAndParseUint(c, "id")
	if valid {
		return
	}

	var request model_actor.RequestUpdateActor
	if valid = pipeline.BindAndValidateRequest(c, ctr.validator, &request); valid {
		return // Error response handled in bindAndValidateRequest
	}

	status, err := ctr.client.UpdateActorById(c, id, request)
	//check status
	if err != nil || !helper.IsSuccessStatus(status) {
		pipeline.AbortWithStatusJSON(c, status, err.Error())
		return
	}
	response := model_actor.ResponseActor{
		Username:  request.Username,
		RoleID:    0,
		Active:    request.Active,
		Verified:  request.Verified,
		UpdatedAt: helper.ConvertTimeToWIB(time.Now()),
	}
	pipeline.JSON(c, status, "update actor", response)
}

func (ctr *ControllerActor) DeleteActorById(c *gin.Context) {
	if valid := pipeline.ValidateJWT(c); valid {
		return // Error response handled in validateJWT
	}

	//bind param
	id, valid := pipeline.BindParamAndParseUint(c, "id")
	if valid {
		return
	}

	status, err := ctr.client.DeleteActorById(c, id)
	//check status
	if err != nil || !helper.IsSuccessStatus(status) {
		pipeline.AbortWithStatusJSON(c, status, err.Error())
		return
	}
	pipeline.JSON(c, status, "Delete Actor", "Success")
}
