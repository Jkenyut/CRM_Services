package helper

import (
	"crm_service/app/model/original"
	"fmt"
	"github.com/go-playground/validator/v10"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"net/http"
	"time"
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
			messageErr += cases.Title(language.Und, cases.NoLower).String(
				fmt.Sprint(messageError.Field(), " must ", messageError.ActualTag(), " ", messageError.Param()))
		}
	}
	return http.StatusPreconditionFailed,
		original.DefaultErrorResponseWithMessage(messageErr, http.StatusPreconditionFailed)
}

func ConvertTimeToWIB(t time.Time) string {
	var format string
	if t.IsZero() {
		return format
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	format = t.In(loc).Format("02-01-2006 15:04:05 MST")
	return format
}
