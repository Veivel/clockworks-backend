package middlewares

import (
	"clockworks-backend/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func jwtKeyFunc(t *jwt.Token) (interface{}, error) {
	return utils.JWT_KEY, nil
}

func JWTMiddleware(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")
	if !strings.Contains(authHeader, " ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid auth header.",
		})
		return
	}

	_, tokenString := strings.Split(authHeader, " ")[0], strings.Split(authHeader, " ")[1]
	// fmt.Println("Auth with", authType, "token")

	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid auth header.",
		})
	}
	// parsing token jwt
	claims := utils.JWTClaim{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, jwtKeyFunc)

	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		switch v.Errors {
		case jwt.ValidationErrorSignatureInvalid:
			// token invalid
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token.",
			})
		case jwt.ValidationErrorExpired:
			// token expired
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Expired token. Please re-login.",
			})
		default:
			c.AbortWithStatus(http.StatusUnauthorized)
			fmt.Println("4")
			return
		}
	}

	if !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// var user models.User
	// models.DB.First()
	// c.Set("User", )
	c.Next()
}
