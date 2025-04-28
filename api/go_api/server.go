package main

import (
	"fmt"
	"net/http"

	"github.com/AdilAnuarbek/akinator-nfactorial/api/go_api/controllers"
	"github.com/AdilAnuarbek/akinator-nfactorial/api/go_api/models"
	"github.com/AdilAnuarbek/akinator-nfactorial/api/go_api/templates"
	"github.com/AdilAnuarbek/akinator-nfactorial/api/go_api/views"
	"github.com/go-chi/chi/v5"
)

var (
	router          *chi.Mux
	handlerServices *models.Akinator
	handlers        *controllers.Handlers
)

func main() {
	router = chi.NewRouter()
	assetsHandler := http.FileServer(http.Dir("api/go_api/public/static"))
	router.Get("/api/go_api/public/static/*", http.StripPrefix("/api/go_api/public/static", assetsHandler).ServeHTTP)
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
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
