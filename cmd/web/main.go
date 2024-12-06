package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kevinchayadi/go-modern-app/pkg/config"
	"github.com/Kevinchayadi/go-modern-app/pkg/handlers"
	"github.com/Kevinchayadi/go-modern-app/pkg/render"
)

var portnumber = ":8080"





func main() {
	var app config.AppConfig

	 tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepository(&app)
	handlers.NewHandler(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	
	// fmt.Println(fmt.Sprintf("Starting application with port %s", portnumber))
	fmt.Printf("Starting application with port %s\n", portnumber)


	_ = http.ListenAndServe(portnumber, nil)
}