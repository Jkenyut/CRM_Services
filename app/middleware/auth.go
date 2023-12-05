package middleware

import (
	"crm_service/app/clients/repository/repository_auth"
	"crm_service/app/config"
	"crm_service/app/model/origin"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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
	var err error
	agent := c.GetHeader("User-Agent")
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		m.AbortWithStatusJSON(c, http.StatusUnauthorized, "Authorization header is missing")
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		m.AbortWithStatusJSON(c, http.StatusUnauthorized, "Invalid authorization header format")
		return
	}

	accessToken := headerParts[1]

	tokenAccess, err := jwt.ParseWithClaims(accessToken, &origin.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.conf.JWT.JwtAccess), nil
	})

	if err.Error() == "token signature is invalid: signature is invalid" {
		m.AbortWithStatusJSON(c, http.StatusUnauthorized, err.Error())
		return
	}

	claimsAccess, ok := tokenAccess.Claims.(*origin.CustomClaims)
	if !ok {
		m.AbortWithStatusJSON(c, http.StatusUnauthorized, "mapping jwt failed")
		return
	}

	externalID := uuid.New().String()
	subject, _ := claimsAccess.GetSubject()
	issuedAt, _ := claimsAccess.GetIssuedAt()
	audience, _ := claimsAccess.GetAudience()
	if len(audience) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, origin.DefaultErrorResponseWithMessage("audience null", http.StatusUnauthorized))
		return
	}

	if claimsAccess.ExpiresAt.Before(time.Now()) {
		var status int
		var JwtRefresh origin.JWTModel

		status, JwtRefresh, err = m.client.CheckSession(c, subject)
		if err != nil || status < 200 || status > 299 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, origin.DefaultErrorResponseWithMessage(err.Error(), http.StatusUnauthorized))
			return
		}

		if issuedAt.Time != JwtRefresh.IssuedAt || audience[1] != agent {
			c.AbortWithStatusJSON(http.StatusUnauthorized, origin.DefaultErrorResponseWithMessage("invalid token data", http.StatusUnauthorized))
			return
		}

		var newTokenAccess string
		status, newTokenAccess, _, err = m.client.GenerateJWTAccessCustom(c, audience[0], audience[1], subject, externalID)

		if err != nil || status < 200 || status > 299 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, origin.DefaultErrorResponseWithMessage("Failed to generate new access token", http.StatusUnauthorized))
			return
		}

		c.Header("Authorization", newTokenAccess)

	} else if audience[1] != agent {
		c.AbortWithStatusJSON(http.StatusUnauthorized, origin.DefaultErrorResponseWithMessage("User agent mismatch", http.StatusUnauthorized))
		return
	}

	c.Set("envJWT", claimsAccess)
	c.Next()
}
