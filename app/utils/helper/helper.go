package helper

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"time"
)

func RequestValidate(err error) (messageErr string) {
	for _, messageError := range err.(validator.ValidationErrors) {
		if messageError.Tag() != "" {
			//message
			messageErr += cases.Title(language.Und, cases.NoLower).String(fmt.Sprint(messageError.Field(), " must ", messageError.ActualTag(), " ", messageError.Param()))
		}
		break
	}
	return messageErr
}

func ConvertTimeToWIB(t time.Time) (format string) {
	if t.IsZero() {
		return format
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	format = t.In(loc).Format("02-01-2006 15:04:05 MST")
	return format
}

func IsSuccessStatus(status int) bool {
	return status >= 200 && status <= 299
}
