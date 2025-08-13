package middleware

import (
	"go-simple-api/response"

	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			response.ErrorResponse(c, response.ErrorCodeTokenInvalid)
			return
		}
		c.Next()
	}
}
