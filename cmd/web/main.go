package main

import (
	"fmt"
	"net/http"

	"github.com/ETISDev/go-bookings/pkg/handlers"
)

const posrt = ":8080"

// main function sets up the HTTP server and routes
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Server started at " + posrt)
	_ = http.ListenAndServe(posrt, nil)
}
