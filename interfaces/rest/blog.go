package rest

import (
	"go-service-gin/application"
	"go-service-gin/util/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BlogREST is
type BlogREST struct {
	blogLogic *application.BlogApplication
}

// NewBlogREST is
func NewBlogREST(ba *application.BlogApplication) *BlogREST {
	if ba == nil {
		return nil
	}

	return &BlogREST{ba}
}

// Fetch is
func (bh *BlogREST) Fetch(c *gin.Context) {
	// bind

	// validation

	// logic
	result, err := bh.blogLogic.Fetch()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	// response
	response.Success(c, http.StatusOK, "success", result)
}
