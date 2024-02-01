package controllers

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/layout/app.html", "views/home.html")
	if err != nil {
		return
	}
	tmpl.Execute(w, nil)
}
func Shop(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/layout/app.html", "views/shop.html")
	if err != nil {
		return
	}
	tmpl.Execute(w, nil)
}
