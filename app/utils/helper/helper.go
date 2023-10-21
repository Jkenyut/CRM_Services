package helper

import (
	"crm_service/app/model/original"
	"fmt"
	"github.com/go-playground/validator/v10"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"net/http"
)

func RequestValidate(err error) (int, original.DefaultResponse) {
	var messageErr string
	firstMessage := true

	for _, messageError := range err.(validator.ValidationErrors) {
		if messageError.Tag() != "" {
			//add ", "
			if !firstMessage {
				messageErr += ", "
			} else {
				firstMessage = false
			}
			//message
			messageErr += cases.Title(language.Und, cases.NoLower).String(fmt.Sprint(messageError.Field(), " must ", messageError.ActualTag(), " ", messageError.Param()))
		}
	}
	return http.StatusPreconditionFailed, original.DefaultErrorResponseWithMessage(messageErr, http.StatusPreconditionFailed)
}
