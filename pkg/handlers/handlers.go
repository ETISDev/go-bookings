package handlers

import (
	"net/http"

	"github.com/ETISDev/go-bookings/pkg/render"
)

// Home and About are HTTP handlers that write a message to the response writer
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.gotpl")

}

// About is an HTTP handler that writes a message to the response writer
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.gotpl")
}
