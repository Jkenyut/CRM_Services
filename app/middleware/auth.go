package middleware

import (
	"crm_service/app/clients/repository/repository_auth"
	"crm_service/app/config"
	"crm_service/app/model/origin"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

type InterfacesMiddlewareAuth interface {
	Auth(c *gin.Context)
}

type AuthMiddleware struct {
	conf   *config.Config
	client repository_auth.InterfaceAuth
}

func NewMiddlewareAuth(conf *config.Config, client repository_auth.InterfaceAuth) InterfacesMiddlewareAuth {
	return &AuthMiddleware{
		conf:   conf,
		client: client,
	}
}

func (m *AuthMiddleware) Auth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, origin.DefaultErrorResponseWithMessage("Authorization header is missing", http.StatusUnauthorized))
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, origin.DefaultErrorResponseWithMessage("Invalid authorization header format", http.StatusUnauthorized))
		return
	}

	accessToken := headerParts[1]
	agent := c.GetHeader("User-Agent")

	claimsAccess, err := m.ParseJWT(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, origin.DefaultErrorResponseWithMessage(err.Error(), http.StatusUnauthorized))
		return
	}

	dataAccess := claimsAccess.Data.(map[string]interface{})
	if claimsAccess.ExpiresAt.Before(time.Now()) {
		var status int
		var JwtRefresh origin.JWTModel
		status, JwtRefresh, err = m.client.CheckSession(c, dataAccess["activity_id"].(string))
		if err != nil || status < 200 || status > 299 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, origin.DefaultErrorResponseWithMessage("Session check failed", http.StatusUnauthorized))
			return
		}
		var claimsRefresh *origin.CustomClaims
		claimsRefresh, err = m.ParseJWT(JwtRefresh.JWTRefresh)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, origin.DefaultErrorResponseWithMessage(err.Error(), http.StatusUnauthorized))
			return
		}

		dataRefresh := claimsRefresh.Data.(map[string]interface{})
		if claimsRefresh.ExpiresAt.Before(time.Now()) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, origin.DefaultErrorResponseWithMessage("Refresh token expired", http.StatusUnauthorized))
			return
		}

		if dataAccess["user_agent"] != agent || dataAccess["activity_id"] != dataRefresh["activity_id"] {
			c.AbortWithStatusJSON(http.StatusUnauthorized, origin.DefaultErrorResponseWithMessage("Invalid token data", http.StatusUnauthorized))
			return
		}
		var newTokenAccess string
		status, newTokenAccess, err = m.client.GenerateJWTAccessCustom(c, dataAccess["role"].(int), agent, dataAccess["activity_id"].(string))
		if err != nil || status < 200 || status > 299 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, origin.DefaultErrorResponseWithMessage("Failed to generate new access token", http.StatusUnauthorized))
			return
		}
		c.Header("Authorization", newTokenAccess)
	} else if dataAccess["user_agent"] != agent {
		c.AbortWithStatusJSON(http.StatusUnauthorized, origin.DefaultErrorResponseWithMessage("User agent mismatch", http.StatusUnauthorized))
		return
	}

	c.Set("envJWT", dataAccess)
	c.Next()
}

func (m *AuthMiddleware) ParseJWT(token string) (*origin.CustomClaims, error) {
	tokenAccess, err := jwt.ParseWithClaims(token, &origin.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.conf.JWT.JwtAccess), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid access token signature: %v", err)
	}

	claims, ok := tokenAccess.Claims.(*origin.CustomClaims)
	if !ok || !tokenAccess.Valid {
		return nil, fmt.Errorf("invalid token claims or token is not valid")
	}
	return claims, nil
}
