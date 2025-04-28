package main

import (
	"net/http"

	"github.com/AdilAnuarbek/akinator-nfactorial/api/go_api/controllers"
	"github.com/AdilAnuarbek/akinator-nfactorial/api/go_api/models"
	"github.com/AdilAnuarbek/akinator-nfactorial/api/go_api/views"
	"github.com/AdilAnuarbek/akinator-nfactorial/public/templates"
	"github.com/go-chi/chi/v5"
)

var (
	router          *chi.Mux
	handlerServices *models.Akinator
	handlers        *controllers.Handlers
)

func init() {
	router = chi.NewRouter()
	assetsHandler := http.FileServer(http.Dir("public/static"))
	router.Get("/public/static/*", http.StripPrefix("/public/static", assetsHandler).ServeHTTP)
	// Home and info pages
	router.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.html", "base.html"))))
	router.Get("/info", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "info.html", "base.html"))))

	// Akinator game pages
	handlerServices = &models.Akinator{Region: "en"}
	handlers = &controllers.Handlers{HandlerServices: handlerServices}
	handlers.Templates.Game = views.Must(views.ParseFS(templates.FS, "akinator-game.html", "base.html"))
	router.Route("/play", func(r chi.Router) {
		r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "akinator-main.html", "base.html"))))
		r.Get("/theme", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "akinator-theme-selection.html", "base.html"))))
		r.Post("/{theme}", handlers.Play)
		r.Post("/answer", handlers.Answer)
		r.Post("/guess", handlers.Guess)
		r.Get("/win", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "akinator-end.html", "base.html"))))
	})

}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
