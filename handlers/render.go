package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func render(w http.ResponseWriter, tmpl string, data any) {
	t, err := template.ParseFiles("templates/layout.html", tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Printf("Erro ao executar o template %s: %v", tmpl, err)
	}
}
