package origin

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	Data any
	jwt.RegisteredClaims
}

type JWTModel struct {
	ActivityId     string `gorm:"column:activity_id"`
	JWTRefresh     string `gorm:"column:jwt_refresh"`
	ExpiredRefresh string `gorm:"column:expired"`
}
