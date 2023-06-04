package middleware

import (
	"crm_service/dto"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage("signature is invalid"))
		c.Abort() // Stop execution of subsequent middleware or handlers
		return
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		if claims.ExpiresAt < time.Now().Unix() {
			c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage("token expired"))
			c.Abort() // Stop execution of subsequent middleware or handlers
			return
		}
		if claims.UserAgent != c.GetHeader("User-Agent") {
			c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage("signature is invalid"))
			c.Abort() // Stop execution of subsequent middleware or handlers
			return
		}
		c.Set("role", claims.Role)
	} else {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage("signature is invalid"))
		c.Abort() // Stop execution of subsequent middleware or handlers
		return
	}

	c.Next()
}
