package main

import (
	"MyFirstApp/pkg/config"
	"MyFirstApp/pkg/handlers"
	"MyFirstApp/pkg/render"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const localHost = "localhost:8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// SessionLoad loads and saves session data for current request
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRep(&app)
	handlers.NewHandles(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    localHost,
		Handler: routers(&app),
	}
	fmt.Println("Server is Started")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
