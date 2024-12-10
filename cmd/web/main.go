package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Kevinchayadi/go-modern-app/pkg/config"
	"github.com/Kevinchayadi/go-modern-app/pkg/handlers"
	"github.com/Kevinchayadi/go-modern-app/pkg/render"
	"github.com/alexedwards/scs/v2"
)

var portnumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager



func main() {
	
	//change to true if production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24*time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	 tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepository(&app)
	handlers.NewHandler(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	
	// fmt.Println(fmt.Sprintf("Starting application with port %s", portnumber))
	fmt.Printf("Starting application with port %s\n", portnumber)
	// _ = http.ListenAndServe(portnumber, nil)

	srv := &http.Server{
		Addr: portnumber,
        Handler: Routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}