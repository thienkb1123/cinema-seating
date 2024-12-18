package utils

import "github.com/gin-gonic/gin"

// ReadBodyRequest body and validate
func ReadBodyRequest(c *gin.Context, request any) error {
	if err := c.ShouldBindJSON(request); err != nil {
		return err
	}

	return ValidateStruct(c.Request.Context(), request)
}
