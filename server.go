package main

import (
	"fmt"
	"net/http"

	"github.com/AdilAnuarbek/akinator-nfactorial/templates"
	"github.com/AdilAnuarbek/akinator-nfactorial/views"
	"github.com/go-chi/chi/v5"
)

type Template interface {
	Execute(w http.ResponseWriter, r *http.Request, data interface{})
}

type Handlers struct {
	Templates struct {
		Index Template
	}
}

func main() {
	r := chi.NewRouter()
	// Home and contact pages
	tmpl, err := views.ParseFS(templates.FS, "home.html", "base.html")
	if err != nil {
		fmt.Printf("parsing template: %v\n", err)
	}
	r.Get("/", StaticHandler(tmpl))

	fmt.Println("Starting the server on 8080...")
	http.ListenAndServe(":8080", r)
}

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, nil)
	}
}
