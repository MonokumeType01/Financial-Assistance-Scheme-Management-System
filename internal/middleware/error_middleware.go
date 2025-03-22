package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			lastError := c.Errors.Last()

			var statusCode = http.StatusInternalServerError
			var errorDetails string

			if meta, ok := lastError.Meta.(string); ok && meta != "" {
				errorDetails = meta
			} else {
				errorDetails = lastError.Error()
			}

			if lastError.Type == gin.ErrorTypePublic {
				statusCode = http.StatusBadRequest
			} else if lastError.Type == gin.ErrorTypeBind {
				statusCode = http.StatusUnprocessableEntity
			} else if lastError.Type == gin.ErrorTypePrivate {
				statusCode = http.StatusNotFound
			}

			c.JSON(statusCode, gin.H{
				"error":   errorDetails,
				"details": lastError.Error(),
			})
			c.Abort()
			return
		}
	}
}
