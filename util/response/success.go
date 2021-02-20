package response

import "github.com/gin-gonic/gin"

// Success is
func Success(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(code, SuccessResponse{
		Error:   false,
		Message: msg,
		Data:    data,
	})
}
