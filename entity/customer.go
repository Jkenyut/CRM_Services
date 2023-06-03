package entity

import "time"

type Customer struct {
	ID        uint      `gorm:"column:id; primaryKey; autoIncrement"`
	FirstName string    `gorm:"column:first_name; index:idx_first_name;size=255"`
	LastName  string    `gorm:"column:last_name; index:idx_last_name;size=255"`
	Email     string    `gorm:"column:email; uniqueIndex"`
	Avatar    string    `gorm:"column:avatar;size:255"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:current_timestamp;autoUpdateTime"`
}

func (Customer) TableName() string {
	return "customer"
}
