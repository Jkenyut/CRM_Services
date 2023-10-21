package model

import "time"

type Customer struct {
	ID        uint      `json:"ID,omitempty" gorm:"column:id; primaryKey; autoIncrement"`
	FirstName string    `json:"firstname ,omitempty" gorm:"column:first_name; index:idx_first_name;size=255"`
	LastName  string    `json:"lastname ,omitempty" gorm:"column:last_name; index:idx_last_name;size=255"`
	Email     string    `json:"email ,omitempty" gorm:"column:email; uniqueIndex"`
	Avatar    string    `json:"avatar ,omitempty" gorm:"column:avatar;size:255"`
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"column:created_at;type:timestamp;original:current_timestamp"`
	UpdatedAt time.Time `json:"updatedAt ,omitempty" gorm:"column:updated_at;type:timestamp;original:current_timestamp;autoUpdateTime"`
	Total     uint64    `json:"total ,omitempty" gorm:"column:total" json:"total,omitempty"`
}

func (Customer) TableName() string {
	return "customer"
}
