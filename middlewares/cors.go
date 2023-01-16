package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// allow CORS
func CORSMiddleware(c *gin.Context) {
	// probably not very safe, but allows for localhost-localhost FE-BE.
	// Edge works just fine though, and I don't know why.
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.Next()
}
