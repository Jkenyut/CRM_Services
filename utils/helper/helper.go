package helper

import (
	"crm_service/dto"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func ValidateData(err error) (int, dto.DefaultResponse) {
	var status = http.StatusPreconditionFailed
	var customErr dto.DefaultResponse
	for _, messageError := range err.(validator.ValidationErrors) {
		messageErr := fmt.Sprint(messageError.StructField(), " ", messageError.ActualTag(), " ", messageError.Param())
		if messageError.Tag() != "" {
			customErr = dto.DefaultErrorResponseWithMessage(messageErr, status)
			break
		}
	}
	return http.StatusPreconditionFailed, customErr
}
