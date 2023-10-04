package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
	"time"
)

func Auth(c *gin.Context) {
	bearer := c.GetHeader("Authorization")
	tokenAuth := strings.Split(bearer, " ")

	var tokenBearer string

	if len(tokenAuth) < 2 {
		// Header is not in the expected format, set a default token value or handle the situation
		tokenBearer = "default_tokenBearer"
	} else {
		tokenBearer = tokenAuth[1]
	}

	token, err := jwt.ParseWithClaims(tokenBearer, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_TOKEN_JWT")), nil
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, "signature token is invalid")
		c.Abort() // Stop execution of subsequent middleware or handlers
		return
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {

		if claims.ExpiresAt.Before(time.Now()) {
			c.JSON(http.StatusBadRequest, "token expired")
			c.Abort() // Stop execution of subsequent middleware or handlers

		}
		if claims.UserAgent != c.GetHeader("User-Agent") {
			c.JSON(http.StatusBadRequest, "signature agent is invalid")
			c.Abort() // Stop execution of subsequent middleware or handlers

		}
		c.Set("role", claims.Role)
	} else {
		c.JSON(http.StatusBadRequest, "signature is invalid")
		c.Abort() // Stop execution of subsequent middleware or handlers

	}

	c.Next()
}
