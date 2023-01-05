package authcontroller

import (
	"clockworks-backend/models"
	"clockworks-backend/utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Index(c *gin.Context) {
	var users []models.User

	models.DB.Find(&users)

	c.JSON(200, gin.H{
		"data": users,
	})
}

func Register(c *gin.Context) {
	var body models.UserData

	err := c.BindJSON(&body)
	if err != nil || body.Email == "" || body.Password == "" || body.Username == "" {
		c.JSON(400, gin.H{
			"message": "Invalid body.",
		})
	} else if utils.GetUser(body.Username).Username != "" {
		c.JSON(400, gin.H{
			"message": "Username already exists.",
		})
	} else {
		enc, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("Error generating a hash with bcrypt.")
			panic(err)
		}

		user := models.User{
			Username:  body.Username,
			Email:     body.Email,
			Password:  string(enc),
			CreatedAt: time.Now(),
		}

		models.DB.FirstOrCreate(&user)

		c.JSON(200, gin.H{
			"data":    user.Username,
			"message": "Succesfully registered user.",
		})
	}
}

func Login(c *gin.Context) {
	var user models.User
	var body models.UserData

	c.BindJSON(&body)

	user = utils.GetUser(body.Username)
	if user.Username == "" {
		c.JSON(400, gin.H{
			"message": "Invalid credentials. (username not found)",
		})
	} else if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid credentials. (incorrect password)",
		})
	} else {
		expTime := time.Now().Add(time.Minute * 10)
		claims := utils.JWTClaim{
			Username: user.Username,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "clockworks-backend",
				ExpiresAt: jwt.NewNumericDate(expTime),
			},
		}

		tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
		token, err := tokenAlgo.SignedString(utils.JWT_KEY)

		if err != nil {
			c.JSON(400, gin.H{
				"message": "Error generating JWT",
			})
		} else {
			c.JSON(200, gin.H{
				"token":   token,
				"message": "Login successful.",
			})
		}
	}
}
