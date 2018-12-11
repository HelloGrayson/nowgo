package main

import (
	"fmt"
	"net/http"
)

// Krupa is a separate lambda for Now
func Krupa(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Krupa on Now!</h1>")
}
