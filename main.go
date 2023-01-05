package main

import (
	"clockworks-backend/controllers/authcontroller"
	"clockworks-backend/controllers/eventcontroller"
	"clockworks-backend/middlewares"
	"clockworks-backend/models"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not found. Will proceed and attempt to use existing environment variables.")
		fmt.Println(os.Getenv("DB_NAME"))
	}

	r := gin.Default()
	models.ConnectDB()

	// Endpoints / Views
	api := r.Group("/api", middlewares.JWTMiddleware)
	api.GET("/events", eventcontroller.Index) // admin
	api.GET("/event/:id", eventcontroller.Show)
	api.POST("/event", eventcontroller.Create)
	api.PATCH("/event/:id", eventcontroller.Update)
	api.DELETE("/event/:id", eventcontroller.Delete)
	// api.GET("/count-events/:id", eventcontroller.CountIds)

	// Authentication Endpoints
	r.GET("/auth/users", authcontroller.Index) // admin
	r.POST("/auth/register", authcontroller.Register)
	r.POST("/auth/login", authcontroller.Login)

	r.Run()
}
