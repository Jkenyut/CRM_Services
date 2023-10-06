package middleware

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	Data any
	jwt.RegisteredClaims
}
