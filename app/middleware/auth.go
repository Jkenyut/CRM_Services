package middleware

import (
	"crm_service/app/clients/repository/repository_auth"
	"crm_service/app/config"
	"crm_service/app/middleware/pipeline"
	"crm_service/app/model/origin"
	"github.com/Jkenyut/libs-numeric-go/libs_auth/libs_auth_jwt"
	"github.com/Jkenyut/libs-numeric-go/libs_models/libs_model_jwt"
	"github.com/Jkenyut/libs-numeric-go/libs_models/libs_model_response"
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
	conf     *config.Config
	client   repository_auth.InterfaceAuth
	libsAuth libs_auth_jwt.InterfacesAuthJWT
}

func NewMiddlewareAuth(conf *config.Config, client repository_auth.InterfaceAuth) InterfacesMiddlewareAuth {
	return &AuthMiddleware{
		conf:     conf,
		client:   client,
		libsAuth: libs_auth_jwt.NewClientAuthJWT(conf.JWT),
	}
}

func (m *AuthMiddleware) Auth(c *gin.Context) {
	var err error
	agent := c.GetHeader("User-Agent")
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		pipeline.AbortWithStatusJSON(c, http.StatusUnauthorized, "Authorization header is missing")
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		pipeline.AbortWithStatusJSON(c, http.StatusUnauthorized, "Invalid authorization header format")
		return
	}

	accessToken := headerParts[1]

	tokenAccess, err := jwt.ParseWithClaims(accessToken, &libs_model_jwt.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.conf.JWT.Access), nil
	})

	if err != nil && err.Error() == "token signature is invalid: signature is invalid" {
		pipeline.AbortWithStatusJSON(c, http.StatusUnauthorized, err.Error())
		return
	}

	claimsAccess, ok := tokenAccess.Claims.(*libs_model_jwt.CustomClaims) // Use pointer type here

	if !ok {
		pipeline.AbortWithStatusJSON(c, http.StatusUnauthorized, "mapping jwt failed")
		return
	}

	externalID := uuid.New().String()
	subject, _ := claimsAccess.GetSubject()
	issuedAt, _ := claimsAccess.GetIssuedAt()
	audience, _ := claimsAccess.GetAudience()
	if len(audience) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, libs_model_response.DefaultErrorResponseWithMessage("audience null", http.StatusUnauthorized))
		return
	}

	ExpiresAt, _ := claimsAccess.GetExpirationTime()
	if ExpiresAt.Before(time.Now()) {
		var status int
		var JwtRefresh origin.JWTModel

		status, JwtRefresh, err = m.client.CheckSession(c, subject)
		if err != nil || status < 200 || status > 299 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, libs_model_response.DefaultErrorResponseWithMessage(err.Error(), http.StatusUnauthorized))
			return
		}

		if issuedAt.Time != JwtRefresh.IssuedAt || audience[1] != agent {
			c.AbortWithStatusJSON(http.StatusUnauthorized, libs_model_response.DefaultErrorResponseWithMessage("invalid token data", http.StatusUnauthorized))
			return
		}

		var newTokenAccess string
		audience = []string{audience[0], audience[1]}
		newTokenAccess, claimsAccess, err = m.libsAuth.GenerateJWTAccessCustom(c, "login", audience, subject, externalID, "")

		if err != nil || status < 200 || status > 299 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, libs_model_response.DefaultErrorResponseWithMessage("Failed to generate new access token", http.StatusUnauthorized))
			return
		}

		c.Header("Authorization", newTokenAccess)
	} else if audience[1] != agent {
		c.AbortWithStatusJSON(http.StatusUnauthorized, libs_model_response.DefaultErrorResponseWithMessage("User agent mismatch", http.StatusUnauthorized))
		return
	}

	c.Set("envJWT", claimsAccess)
	c.Next()
}
