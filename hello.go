package main

import (
	"fmt"
	"github.com/chris-carpenter/hello/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(morestrings.ReverseRunes("Hello, Go!"))
}