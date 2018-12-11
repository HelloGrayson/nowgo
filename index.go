package main

import (
	"fmt"
	"net/http"
)

// Handler is a lambda for Now
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go on Now!</h1>")
}
