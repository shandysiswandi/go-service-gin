package rest

import (
	"go-service-gin/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BlogREST is
type BlogREST struct {
	blogLogic *application.BlogApplication
}

// NewBlogREST is
func NewBlogREST(ba *application.BlogApplication) *BlogREST {
	return &BlogREST{ba}
}

// Fetch is
func (bh *BlogREST) Fetch(c *gin.Context) {
	// costum context

	// bind

	// validation

	// logic

	// response
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
