package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[Error] Panic: %v", err)
				c.JSON(500, gin.H{"error": "Internal server error"})
			}
		}()

		c.Next()

		// Handle errors 
		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			if err.Err == gorm.ErrRecordNotFound {
				c.JSON(404, gin.H{"error": "Record not found"})
				return
			}

			log.Printf("[ERROR] %v", err.Err)
			c.JSON(500, gin.H{"error": "Internal server error"})
		}
	}
}