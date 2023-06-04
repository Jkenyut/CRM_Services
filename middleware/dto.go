package middleware

import "github.com/dgrijalva/jwt-go"

type Response struct {
	Data []User `json:"data"`
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type CustomClaims struct {
	Role      uint   `json:"role"`
	UserAgent string `json:"user_agent"`
	jwt.StandardClaims
}
