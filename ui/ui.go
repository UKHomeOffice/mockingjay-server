package ui

import (
	"html/template"
	"log"
	"net/http"
)

type Handler struct {
	layout *template.Template
}

func NewHandler() *Handler {
	h := new(Handler)

	t, err := template.ParseFiles("ui/template.html")

	if err != nil {
		log.Fatal("Unable to load template file for UI ", err)
	}

	h.layout = t

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.layout.Execute(w, nil)
}
