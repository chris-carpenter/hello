// Package messages implements ouput text
package messages

import (
	"net/http"

	"strings"

	"github.com/chris-carpenter/hello/morestrings"
	"github.com/google/go-cmp/cmp"
	"github.com/gin-gonic/gin"
)

//Hello prints hello message
func Hello(c *gin.Context) {
	hello := []string{ morestrings.ReverseRunes("!oG ,olleH"), cmp.Diff("Hello World", "Hello Go") }
	c.String(http.StatusOK, strings.Join(hello, "\n"))
}