package entity

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strconv"
)

type DefaultResponse struct {
	Message ResponseMeta `json:"message,omitempty"`
	Data    any          `json:"data,omitempty"`
}

func DefaultErrorResponseWithMessage(msg string, status int) DefaultResponse {
	return DefaultResponse{
		Message: ResponseMeta{
			Success:    false,
			Message:    cases.Title(language.Und, cases.NoLower).String(msg),
			StatusCode: strconv.Itoa(status),
		},
	}
}

func DefaultSuccessResponseWithMessage(msg string, status int, data any) DefaultResponse {
	return DefaultResponse{
		Message: ResponseMeta{
			Success:    true,
			Message:    cases.Title(language.Und, cases.NoLower).String(msg),
			StatusCode: strconv.Itoa(status),
		},
		Data: data,
	}
}
