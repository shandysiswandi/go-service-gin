package rest

import (
	"go-service-gin/util/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DefaultREST is
type DefaultREST struct{}

// NewDefaultREST is
func NewDefaultREST() *DefaultREST {
	return &DefaultREST{}
}

// Home is
func (DefaultREST) Home(c *gin.Context) {
	response.Success(c, http.StatusOK, "welcome to home", []string{})
}

// RouteNotFound is
func (DefaultREST) RouteNotFound(c *gin.Context) {
	response.ErrorNotFound(c, "Route Not Found", nil)
}

// MethodNotFound is
func (DefaultREST) MethodNotFound(c *gin.Context) {
	response.ErrorMethodNotAllowed(c)
}
