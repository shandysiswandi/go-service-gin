package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeREST is
type HomeREST struct{}

// NewHomeREST is
func NewHomeREST() *HomeREST {
	return &HomeREST{}
}

// Home is
func (HomeREST) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "welcome to home",
	})
}
