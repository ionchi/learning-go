package main

import (
	"learning/go/controllers"
	"net/http"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil) //nolint:errcheck
}
