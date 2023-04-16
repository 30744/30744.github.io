package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Handler function for the homepage
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Serve the index.html file
		http.ServeFile(w, r, "static/form.html")
	})

	// Handler function for the about page
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		// Serve the about.html file
		http.ServeFile(w, r, "static/about.html")
	})

	// Handler function for the form submission
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		// Parse the form data
		name := r.FormValue("name")
		email := r.FormValue("email")

		// Send a response
		fmt.Fprintf(w, "Thank you for registering, %s! We will contact you at %s.", name, email)
	})

	// Start the server
	fmt.Println("Server started on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
