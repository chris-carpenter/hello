// Package messages implements ouput text
package messages

import (
	"fmt"
	"net/http"

	"strings"

	"github.com/chris-carpenter/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

//Hello prints hello message
func Hello(w http.ResponseWriter, req *http.Request) {
	hello := []string{ morestrings.ReverseRunes("!oG ,olleH"), cmp.Diff("Hello World", "Hello Go") }
	fmt.Fprintln(w, strings.Join(hello, "\n"))
}