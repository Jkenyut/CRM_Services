package entity

import "time"

type Actor struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	Username  string    `gorm:"column:username;unique;size:255"`
	Password  string    `json:"password,omitempty" gorm:"column:password;size:255"`
	RoleID    uint32    `gorm:"column:role_id;default:2"`
	Verified  string    `gorm:"column:verified;type:enum('true','false');default:'false'"`
	Active    string    `gorm:"column:active;type:enum('true','false');default:'false'"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:current_timestamp;autoUpdateTime"`
}
