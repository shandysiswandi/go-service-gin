package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorNotFound is
func ErrorNotFound(c *gin.Context, msg string, err error) {
	c.JSON(http.StatusNotFound, ErrorResponse{
		Error:   true,
		Message: msg,
		Trace:   err,
	})
}

// ErrorMethodNotAllowed is
func ErrorMethodNotAllowed(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, ErrorResponse{
		Error:   true,
		Message: "Method Not Allowed",
		Trace:   nil,
	})
}
