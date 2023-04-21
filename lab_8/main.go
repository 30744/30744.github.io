package main

import (
	"html/template"
	"log"
	"net/http"
)

type Users struct {
	Login         string
	Password      string
	Success       bool
	StorageAccess string
}

var (
	tmpl = template.Must(template.ParseFiles("form.html"))
)

func handler(w http.ResponseWriter, r *http.Request) {
	data := Users{
		Login:    r.FormValue("login"),
		Password: r.FormValue("password"),
	}
	data.Success = true
	data.StorageAccess = "Hello!"
	tmpl.Execute(w, data)

}
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
