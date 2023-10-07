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

type UpdateCustomer struct {
	FirstName string `json:"firstname" validate:"required,min=1,max=100,alpha"`
	LastName  string `json:"lastname" validate:"min=1,max=100,alpha"`
	Avatar    string `json:"avatar" validate:"min=1,max=250,alphanumunicode"`
}

type FindAllCustomer struct {
	Page       uint             `json:"page,omitempty"`
	PerPage    uint             `json:"per_page,omitempty"`
	TotalPages uint             `json:"total_pages,omitempty"`
	Data       []model.Customer `json:"data"`
}
