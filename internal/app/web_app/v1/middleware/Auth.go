package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticateMiddleware(c *gin.Context) {
	// Perform middleware operations here
	_, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "401 UnAuthenticate !!",
			"status":  401,
		})
		return
	}
	// Call the next handler
	c.Next()
}
