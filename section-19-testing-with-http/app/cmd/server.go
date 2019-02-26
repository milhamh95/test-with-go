package main

import (
	"net/http"
	"test-with-go/section-19-testing-with-http/app"
)

func main() {
	http.ListenAndServe(":3000", &app.Server{})
}
