package main

import (
	"fmt"
	"net/http"

	"github.com/AdilAnuarbek/akinator-nfactorial/controllers"
	"github.com/AdilAnuarbek/akinator-nfactorial/models"
	"github.com/AdilAnuarbek/akinator-nfactorial/templates"
	"github.com/AdilAnuarbek/akinator-nfactorial/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	assetsHandler := http.FileServer(http.Dir("static"))
	r.Get("/static/*", http.StripPrefix("/static", assetsHandler).ServeHTTP)
	// Home and info pages
	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.html", "base.html"))))
	r.Get("/info", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "info.html", "base.html"))))

	// Akinator game pages
	handlerServices := &models.Akinator{Region: "en"}
	handlers := controllers.Handlers{HandlerServices: handlerServices}
	handlers.Templates.Game = views.Must(views.ParseFS(templates.FS, "akinator-game.html", "base.html"))
	r.Route("/play", func(r chi.Router) {
		r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "akinator-main.html", "base.html"))))
		r.Get("/theme", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "akinator-theme-selection.html", "base.html"))))
		r.Post("/{theme}", handlers.Play)
		r.Post("/answer", handlers.Answer)
		r.Post("/guess", handlers.Guess)
		r.Get("/win", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "akinator-end.html", "base.html"))))
	})

	fmt.Println("Starting the server on 8080...")
	http.ListenAndServe(":8080", r)
}
