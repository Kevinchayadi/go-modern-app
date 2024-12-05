package handlers

import (
	"net/http"

	"github.com/Kevinchayadi/go-modern-app/pkg/render"
)

func Home(w http.ResponseWriter,r *http.Request){
// fmt.Fprintf(w, "this is the home page")
	render.RenderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
// sum:= AddValue(2,2)
// _, _= fmt.Fprintf(w, fmt.Sprintf("this is the about page, two plus two is %d",sum))
	render.RenderTemplate(w, "about.page.html")
}

