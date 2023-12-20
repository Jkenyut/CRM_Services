package origin

import (
	"time"
)

type JWTModel struct {
	ActivityId string    `gorm:"column:activity_id"`
	Agent      string    `gorm:"column:agent"`
	IssuedAt   time.Time `gorm:"column:issued_at"`
	ExpiredAt  time.Time `gorm:"column:expired_at"`
}
