package dto

import "strings"

type DefaultResponse struct {
	Message ResponseMeta `json:"message,omitempty"`
	Data    any          `json:"data,omitempty"`
}

func DefaultErrorResponseWithMessage(msg string, time string, status string) DefaultResponse {
	return DefaultResponse{
		Message: ResponseMeta{
			Success:      false,
			Message:      strings.Title(msg),
			ResponseTime: time,
			StatusCode:   status,
		},
	}
}

func DefaultSuccessResponseWithMessage(msg string, time string, status string, data any) DefaultResponse {
	return DefaultResponse{
		Message: ResponseMeta{
			Success:      true,
			Message:      strings.Title(msg),
			ResponseTime: time,
			StatusCode:   status,
		},
		Data: data,
	}
}
