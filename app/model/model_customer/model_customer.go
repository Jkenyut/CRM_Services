package model_customer

import "time"

type Customer struct {
	ID        uint64    `json:"ID,omitempty" gorm:"column:id; primaryKey; autoIncrement"`
	FirstName string    `json:"firstname ,omitempty" gorm:"column:firstname; index:idx_first_name;size=255"`
	LastName  string    `json:"lastname ,omitempty" gorm:"column:lastname; index:idx_last_name;size=255"`
	Email     string    `json:"email ,omitempty" gorm:"column:email; uniqueIndex"`
	Avatar    string    `json:"avatar ,omitempty" gorm:"column:avatar;size:255"`
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"column:created_at;type:timestamp;origin:current_timestamp"`
	UpdatedAt time.Time `json:"updatedAt ,omitempty" gorm:"column:updated_at;type:timestamp;origin:current_timestamp;autoUpdateTime"`
	Total     uint64    `json:"total ,omitempty" gorm:"column:total" json:"total,omitempty"`
}

func (Customer) TableName() string {
	return "customers"
}

type RequestCustomer struct {
	FirstName string `json:"firstname,omitempty" validate:"required,min=1,max=100,alpha"`
	LastName  string `json:"lastname,omitempty" validate:"min=1,max=100,alpha"`
	Email     string `json:"email,omitempty" validate:"required,email"`
	Avatar    string `json:"avatar,omitempty" validate:"max=500"`
}

type RequestUpdateCustomer struct {
	FirstName string `json:"firstname,omitempty" validate:"required,min=1,max=100,alpha"`
	LastName  string `json:"lastname,omitempty" validate:"min=1,max=100,alpha"`
	Avatar    string `json:"avatar,omitempty" validate:"max=500"`
}

type FindAllCustomer struct {
	Page       uint64             `json:"page,omitempty"`
	PerPage    uint64             `json:"per_page,omitempty"`
	TotalPages float64            `json:"total_pages,omitempty"`
	TotalData  uint64             `json:"totalData,omitempty"`
	Data       []ResponseCustomer `json:"data"`
}
type RequestCustomerEmail struct {
	Email string `json:"email" validate:"required,email"`
}

type ResponseCustomer struct {
	ID        uint64 `json:"ID,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty" `
	Avatar    string `json:"avatar,omitempty"  `
	CreatedAt string `json:"createdAt,omitempty" gorm:"column:created_at;type:timestamp;default:current_timestamp" `
	UpdatedAt string `json:"updatedAt,omitempty" gorm:"column:updated_at;type:timestamp;default:current_timestamp;autoUpdateTime"`
	Total     uint64 `json:"total,omitempty" `
}
