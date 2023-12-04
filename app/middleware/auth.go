package middleware

import (
	"crm_service/app/config"
	"crm_service/app/model/origin"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

func Auth(c *gin.Context) {
	bearerAccess := c.GetHeader("Authorization")
	bearerRefresh := c.GetHeader("Refresh-Token")
	tokenAccessAuth := strings.Split(bearerAccess, " ")
	tokenRefreshAuth := strings.Split(bearerRefresh, " ")

	var tokenAccessBearer, tokenRefreshBearer string

	if len(tokenRefreshAuth) < 2 {
		// Header is not in the expected format, set a origin token value or handle the situation
		tokenRefreshBearer = "default_tokenBearer"
	} else {
		tokenRefreshBearer = tokenRefreshAuth[1]
	}

	if len(tokenAccessAuth) < 2 {
		// Header is not in the expected format, set a origin token value or handle the situation
		tokenAccessBearer = "default_tokenBearer"
	} else {
		tokenAccessBearer = tokenAccessAuth[1]
	}

	tokenRefresh, err := jwt.ParseWithClaims(tokenRefreshBearer, &origin.CustomClaims{}, func(tokenRefresh *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().JWT.JwtRefresh), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, origin.DefaultErrorResponseWithMessage("signature token is invalid", http.StatusUnauthorized)) // Stop execution of subsequent middleware or handlers
	}
	tokenAccess, err := jwt.ParseWithClaims(tokenAccessBearer, &origin.CustomClaims{}, func(tokenAccess *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().JWT.JwtAccess), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, origin.DefaultErrorResponseWithMessage("signature token is invalid", http.StatusUnauthorized)) // Stop execution of subsequent middleware or handlers

	}
	claimsRefresh := tokenRefresh.Claims.(*origin.CustomClaims)
	if tokenRefresh.Valid {
		data := claimsRefresh.Data.(map[string]interface{})
		if claimsRefresh.ExpiresAt.Before(time.Now()) {
			c.AbortWithStatusJSON(http.StatusBadRequest, origin.DefaultErrorResponseWithMessage("token refresh expired", http.StatusUnauthorized)) // Stop execution of subsequent middleware or handlers
		}
		if data["user_agent"] != c.GetHeader("User-Agent") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, origin.DefaultErrorResponseWithMessage("signature agent is invalid", http.StatusUnauthorized)) // Stop execution of subsequent middleware or handlers
		}
	}

	if claimsAccess, ok := tokenAccess.Claims.(*origin.CustomClaims); ok && tokenAccess.Valid {
		data := claimsAccess.Data.(map[string]interface{})
		if claimsAccess.ExpiresAt.Before(time.Now()) {
			claimsRefresh
			c.AbortWithStatusJSON(http.StatusBadRequest, origin.DefaultErrorResponseWithMessage("token access expired", http.StatusUnauthorized)) // Stop execution of subsequent middleware or handlers
		}
		if data["user_agent"] != c.GetHeader("User-Agent") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, origin.DefaultErrorResponseWithMessage("signature agent is invalid", http.StatusUnauthorized)) // Stop execution of subsequent middleware or handlers

		}
		c.Set("envJWT", data)
	}

	c.Next()
}
