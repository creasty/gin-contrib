package frontend

import (
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	funcName = "/creasty/gin-contrib/frontend.Wrap"
)

// Wrap is a middleware that fallbacks to the given handler
// when no handler for a path is defined in the router
func Wrap(fn gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "GET" && c.Request.Method != "HEAD" {
			c.Next()
			return
		}

		c.Next()

		if strings.Contains(c.HandlerName(), funcName) {
			fn(c)
		}
	}
}
