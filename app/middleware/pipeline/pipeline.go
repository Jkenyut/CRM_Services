package pipeline

import (
	"crm_service/app/utils/helper"
	"github.com/Jkenyut/libs-numeric-go/libs_models/libs_model_jwt"
	"github.com/Jkenyut/libs-numeric-go/libs_models/libs_model_response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

func AbortWithStatusJSON(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, libs_model_response.DefaultErrorResponseWithMessage(message, status))
}

func JSON(c *gin.Context, status int, message string, data any) {
	c.JSON(status, libs_model_response.DefaultSuccessResponseWithMessage(message, status, data))
}

func BindAndValidateRequest(c *gin.Context, v *validator.Validate, request interface{}) bool {
	if err := c.BindJSON(request); err != nil {
		AbortWithStatusJSON(c, http.StatusBadRequest, "required not valid")
		return true
	}

	if err := v.Struct(request); err != nil {
		AbortWithStatusJSON(c, http.StatusPreconditionFailed, helper.RequestValidate(err))
		return true
	}
	return false
}

func ValidateJWT(c *gin.Context) bool {
	envJWT, ok := c.Get("envJWT")
	if !ok {
		AbortWithStatusJSON(c, http.StatusForbidden, "env jwt not found")
		return true
	}
	setJWT := envJWT.(*libs_model_jwt.CustomClaims)
	audience, _ := setJWT.GetAudience()
	if len(audience) == 0 || audience[0] != "1" {
		AbortWithStatusJSON(c, http.StatusUnauthorized, "Account Not Authorized")
		return true
	}
	return false
}

func BindParamAndParseUint(c *gin.Context, request string) (id uint64, valid bool) {
	res, err := strconv.ParseUint(c.Param(request), 10, 64)
	if err != nil {
		AbortWithStatusJSON(c, http.StatusBadRequest, "must unsigned number")
		return 0, true
	}
	return res, false
}

func BindQueryAndParseUint(c *gin.Context, request string, defaultValue string) (id uint64, valid bool) {
	res, err := strconv.ParseUint(c.DefaultQuery(request, defaultValue), 10, 64)
	if err != nil {
		AbortWithStatusJSON(c, http.StatusBadRequest, "must unsigned number")
		return
	}
	return res, false
}
