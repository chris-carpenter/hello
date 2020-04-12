// Package messages implements ouput text
package messages

import (
	"github.com/spf13/viper"
	"net/http"

	"strings"

	"github.com/chris-carpenter/hello/morestrings"
	"github.com/google/go-cmp/cmp"
	"github.com/gin-gonic/gin"
)

//Hello prints hello message
func Hello(c *gin.Context) {
	message := viper.GetString("helloMessage")
	hello := []string{ morestrings.ReverseRunes(message), cmp.Diff("Hello World", "Hello Go") }
	c.String(http.StatusOK, strings.Join(hello, "\n"))
}