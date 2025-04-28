package controllers

import (
	"net/http"

	"github.com/AdilAnuarbek/akinator-nfactorial/api/go_api/models"
	"github.com/go-chi/chi/v5"
)

type Template interface {
	Execute(w http.ResponseWriter, r *http.Request, data any)
}

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, nil)
	}
}

type Handlers struct {
	Templates struct {
		Game Template
	}
	HandlerServices *models.Akinator
}

// enum in node_akinator code:
// Character = 1
// Animals = 14
// Objects = 2
func (h Handlers) Play(w http.ResponseWriter, r *http.Request) { // POST
	theme := chi.URLParam(r, "theme")
	h.HandlerServices.Theme = theme
	question, err := h.HandlerServices.StartAkinatorGame()
	if err != nil {
		http.Error(w, "Failed to start Akinator game", http.StatusInternalServerError)
		return
	}
	data := struct {
		Win      bool
		Name     string
		Question string
	}{
		Win:      false,
		Name:     "",
		Question: question,
	}
	h.Templates.Game.Execute(w, r, data)
}

func (h Handlers) Answer(w http.ResponseWriter, r *http.Request) { // POST
	answer := r.FormValue("answer") // 0-4
	question, win, name, err := h.HandlerServices.AnswerAkinatorGame(answer)
	if err != nil {
		http.Error(w, "Failed to answer Akinator game", http.StatusInternalServerError)
		return
	}
	data := struct {
		Win      bool
		Name     string
		Question string
		Progress uint8
	}{
		Win:      win,
		Name:     name,
		Question: question,
	}
	h.Templates.Game.Execute(w, r, data)
}

func (h Handlers) Guess(w http.ResponseWriter, r *http.Request) { // POST
	guess := r.FormValue("guess") // 1 - true, 0 - false
	question, err := h.HandlerServices.GuessAkinatorGame(guess)
	if err != nil {
		http.Error(w, "Failed to answer Akinator game", http.StatusInternalServerError)
		return
	}
	if guess == "1" {
		http.Redirect(w, r, "/play/win", http.StatusFound)
		return
	}
	if guess == "0" {
		data := struct {
			Win      bool
			Name     string
			Question string
		}{
			Win:      false,
			Name:     "",
			Question: question,
		}
		h.Templates.Game.Execute(w, r, data)
	}
}
