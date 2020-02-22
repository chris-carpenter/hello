// Package health implements healthcheck mechanisms
package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Headers shows header
func Headers(c *gin.Context) {
	mHeaders := make(map[string]string)
	for name, headers := range c.Request.Header {
		for _, h := range headers {
			mHeaders[name]=h
		}
	}
	c.JSON(http.StatusOK, mHeaders)
}