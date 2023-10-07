package customer

import (
	"crm_service/entity"
	"crm_service/model"
)

type CustomerBody struct {
	FirstName string `json:"firstname" validate:"required,min=1,max=100,alpha"`
	LastName  string `json:"lastname" validate:"min=1,max=100,alpha"`
	Email     string `json:"email" validate:"required,email"`
	Avatar    string `json:"avatar" validate:"min=1,max=250,alphanumunicode"`
}

type UpdateCustomerBody struct {
	FirstName string `json:"firstname" validate:"required,min=1,max=100,alpha"`
	LastName  string `json:"lastname" validate:"min=1,max=100,alpha"`
	Avatar    string `json:"avatar" validate:"min=1,max=250,alphanumunicode"`
}

type SuccessCreate struct {
	entity.ResponseMeta
	Data CustomerBody `json:"data"`
}

type FindCustomer struct {
	entity.ResponseMeta
	Data model.Customer `json:"data"`
}

type FindAllCustomer struct {
	entity.ResponseMeta
	Page       uint             `json:"page,omitempty"`
	PerPage    uint             `json:"per_page,omitempty"`
	Total      int              `json:"total,omitempty"`
	TotalPages uint             `json:"total_pages,omitempty"`
	Data       []model.Customer `json:"data"`
}
