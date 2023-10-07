package entity

import (
	"strconv"
	"strings"
)

type DefaultResponse struct {
	Message ResponseMeta `json:"message,omitempty"`
	Data    any          `json:"data,omitempty"`
}

func DefaultErrorResponseWithMessage(msg string, status int) DefaultResponse {
	return DefaultResponse{
		Message: ResponseMeta{
			Success:    false,
			Message:    strings.Title(msg),
			StatusCode: strconv.Itoa(status),
		},
	}
}

func DefaultSuccessResponseWithMessage(msg string, status int, data any) DefaultResponse {
	return DefaultResponse{
		Message: ResponseMeta{
			Success:    true,
			Message:    strings.Title(msg),
			StatusCode: strconv.Itoa(status),
		},
		Data: data,
	}
}
