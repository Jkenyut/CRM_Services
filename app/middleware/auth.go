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
	    // Extract the Authorization header
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
        c.AbortWithStatusJSON(http.StatusUnauthorized, origin.DefaultErrorResponseWithMessage("Authorization header is missing", http.StatusUnauthorized))
        return
    }

    // Split the header into 'Bearer' and the token
    headerParts := strings.Split(authHeader, " ")
    if len(headerParts) != 2 || headerParts[0] != "Bearer" {
        c.AbortWithStatusJSON(http.StatusUnauthorized, origin.DefaultErrorResponseWithMessage("Invalid authorization header format", http.StatusUnauthorized))
        return
    }

    accessToken := headerParts[1]

	claimsAccess, err := m.ParseJWT(accessToken)
	dataAccess := claimsAccess.Data.(map[string]interface{})

	if dataAccess["user_agent"] != c.GetHeader("User-Agent") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, origin.DefaultErrorResponseWithMessage("signature agent is invalid", http.StatusUnauthorized)) // Stop execution of subsequent middleware or handlers
	}
		c.Set("envJWT", dataAccess)
	}
	c.Next()
}

func (m *AuthMiddleware) ParseJWT(token string) (claims *origin.CustomClaims,err error){
	    // Parse the access token
    tokenAccess, err := jwt.ParseWithClaims(token, &origin.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(m.conf.JWT.JwtAccess), nil
    })
    if err != nil {
        return claims,fmt.Errorf("Invalid access token signature")
    }

    // Validate the access token
    claimsAccess, ok := tokenAccess.Claims.(*origin.CustomClaims)
    if !ok || !tokenAccess.Valid || claimsAccess.ExpiresAt.Before(time.Now()) {
       return  claims,fmt.Errorf("Invalid or expired access token")
    }
	return claimsAccess,nil
}


