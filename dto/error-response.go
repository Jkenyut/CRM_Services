package dto

type ErrorResponse struct {
	ResponseMeta
	Data   any `json:"data"`
	Errors any `json:"errors,omitempty"`
}

func DefaultErrorResponse() ErrorResponse {
	return DefaultErrorResponseWithMessage("")
}

func DefaultErrorResponseWithMessage(msg string) ErrorResponse {
	return ErrorResponse{
		ResponseMeta: ResponseMeta{
			Success:      false,
			MessageTitle: "Oops, something went wrong.",
			Message:      msg,
			ResponseTime: "",
		},
		Data: nil,
	}
}

func DefaultErrorInvalidDataWithMessage(msg string) ErrorResponse {
	return ErrorResponse{
		ResponseMeta: ResponseMeta{
			Success:      false,
			MessageTitle: "Oops, something went wrong.",
			Message:      "Form Invalid data.",
			ResponseTime: "",
		},
		Data: msg,
	}
}

func DefaultDataInvalidResponse(validationErrors any) ErrorResponse {
	return ErrorResponse{
		ResponseMeta: ResponseMeta{
			MessageTitle: "Oops, something went wrong.",
			Message:      "Data invalid.",
		},
		Errors: validationErrors,
	}
}

func DefaultBadRequestResponse() ErrorResponse {
	return DefaultErrorResponseWithMessage("Bad request")
}
