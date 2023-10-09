package customer

import (
	"crm_service/model"
)

type RequestCustomer struct {
	FirstName string `json:"firstname" validate:"required,min=1,max=100,alpha"`
	LastName  string `json:"lastname" validate:"min=1,max=100,alpha"`
	Email     string `json:"email" validate:"required,email"`
	Avatar    string `json:"avatar" validate:"min=1,max=250,alphanumunicode"`
}

type RequestUpdateCustomer struct {
	FirstName string `json:"firstname" validate:"required,min=1,max=100,alpha"`
	LastName  string `json:"lastname" validate:"min=1,max=100,alpha"`
	Avatar    string `json:"avatar" validate:"min=1,max=250,alphanumunicode"`
}

type FindAllCustomer struct {
	Page       uint64           `json:"page,omitempty"`
	PerPage    uint64           `json:"per_page,omitempty"`
	TotalPages uint64           `json:"total_pages,omitempty"`
	Data       []model.Customer `json:"data"`
}

type RequestGetAllCustomer struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type RequestBulkCustomer struct {
	FirstName string `json:"first_name" validate:"required,min=1,max=100,alpha"`
	LastName  string `json:"last_name" validate:"min=1,max=100,alpha"`
	Email     string `json:"email" validate:"required,email"`
	Avatar    string `json:"avatar" validate:"min=1,max=250,alphanumunicode"`
}
type ResponseBulkData struct {
	Data []RequestBulkCustomer `json:"data"`
}
