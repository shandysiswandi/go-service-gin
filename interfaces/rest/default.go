package rest

import (
	"go-service-gin/infrastructure/library/sentry"
	"go-service-gin/util/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DefaultREST is
type DefaultREST struct {
	sentry *sentry.Sentry
}

// NewDefaultREST is
func NewDefaultREST(s *sentry.Sentry) *DefaultREST {
	return &DefaultREST{s}
}

// Home is
func (DefaultREST) Home(c *gin.Context) {
	response.Success(c, http.StatusOK, "welcome to home", []string{})
}

// Sentry is
func (d *DefaultREST) Sentry(c *gin.Context) {
	e := d.sentry.Message("sentry check")
	response.Success(c, http.StatusOK, "sentry check | if data null, sentry not connect", e)
}

// RouteNotFound is
func (DefaultREST) RouteNotFound(c *gin.Context) {
	response.ErrorNotFound(c, "Route Not Found", nil)
}

// MethodNotFound is
func (DefaultREST) MethodNotFound(c *gin.Context) {
	response.ErrorMethodNotAllowed(c)
}
