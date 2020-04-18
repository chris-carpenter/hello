// Package messages implements ouput text
package messages

import (
	"encoding/json"
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
	settings, err := json.Marshal(viper.AllSettings())
	if  err !=nil {
		panic(err)
	}
	//id := "battle_client_id: " + string()
	hello := []string{ morestrings.ReverseRunes(message), cmp.Diff("Hello World", "Hello Go"), string(settings) }
	c.String(http.StatusOK, strings.Join(hello, "\n"))
}