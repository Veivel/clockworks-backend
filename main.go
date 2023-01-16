package main

import (
	"clockworks-backend/controllers/authcontroller"
	"clockworks-backend/controllers/eventcontroller"
	"clockworks-backend/controllers/tagcontroller"
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
	def := gin.Default()
	r := def.Group("/", middlewares.CORSMiddleware)
	models.ConnectDB()

	// Event Endpoints (req AUTH)

	api := r.Group("/api", middlewares.JWTMiddleware)
	api.GET("/events", eventcontroller.Index)
	api.POST("/event", eventcontroller.Create)
	api.GET("/event/:id", eventcontroller.Show)
	api.PATCH("/event/:id", eventcontroller.Update)
	api.DELETE("/event/:id", eventcontroller.Delete)

	// Tag Endpoints (req AUTH)

	api.GET("/tags/:eventId", tagcontroller.Index)
	api.POST("/tags/:eventId", tagcontroller.Toggle)

	// Authentication Endpoints

	api.GET("/profile", authcontroller.Profile) // requires auth, hence api group
	r.POST("/auth/register", authcontroller.Register)
	r.POST("/auth/login", authcontroller.Login)

	// Admin-only endpoints (TODO: implement special auth, with different user model)

	admin := r.Group("/admin")
	admin.GET("/events", eventcontroller.All)
	admin.GET("/users", authcontroller.All)

	def.Run()
}
