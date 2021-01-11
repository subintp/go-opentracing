package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

// SimpleMiddleware - simple test middleware
func SimpleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Logging from simple middleware")
		c.Next()
	}
}
