package main

import (
	"clockworks-backend/controllers/eventcontroller"
	"clockworks-backend/models"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not found. Will proceed and attempt to use environment variables.")
		fmt.Println(os.Getenv("DB_NAME"))
	}

	r := gin.Default()
	models.ConnectDB()

	// Endpoints / Views
	r.GET("/api/events", eventcontroller.Index)
	r.GET("/api/count-events/:id", eventcontroller.CountIds)
	r.GET("/api/event/:id", eventcontroller.Show)
	r.POST("/api/event", eventcontroller.Create)
	r.PATCH("/api/event/:id", eventcontroller.Update)
	r.DELETE("/api/event/:id", eventcontroller.Delete)

	r.Run()
}
