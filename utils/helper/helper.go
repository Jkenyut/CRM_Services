package helper

import (
	"crm_service/entity"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func ValidateData(err error) (int, entity.DefaultResponse) {
	var status = http.StatusPreconditionFailed
	var customErr entity.DefaultResponse
	for _, messageError := range err.(validator.ValidationErrors) {
		messageErr := fmt.Sprint(messageError.Field(), " ", messageError.ActualTag(), " ", messageError.Param())
		if messageError.Tag() != "" {
			customErr = entity.DefaultErrorResponseWithMessage(messageErr, status)
			break
		}
	}
	return http.StatusPreconditionFailed, customErr
}
