package messages

import (

	"github.com/chris-carpenter/hello/morestrings"
	"github.com/google/go-cmp/cmp"
	"strings"
)

//Hello prints hello message
func Hello() string {
	hello := []string{ morestrings.ReverseRunes("!oG ,olleH"), cmp.Diff("Hello World", "Hello Go") }
	return strings.Join(hello, "\n")
}