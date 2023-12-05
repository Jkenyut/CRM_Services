package origin

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomClaims struct {
	Data any
	jwt.RegisteredClaims
}

type JWTModel struct {
	ActivityId string    `gorm:"column:activity_id"`
	Agent      string    `gorm:"column:agent"`
	IssuedAt   time.Time `gorm:"column:issued_at"`
	ExpiredAt  time.Time `gorm:"column:expired_at"`
}
