package main

import (
	"net/http"

	"github.com/chris-carpenter/hello/messages"
	"github.com/chris-carpenter/hello/health"
)

func main() {
	http.HandleFunc("/hello", messages.Hello)
	http.HandleFunc("/headers", health.Headers)

	http.ListenAndServe(":8090", nil)
}