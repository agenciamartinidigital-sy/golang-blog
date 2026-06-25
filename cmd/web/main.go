package main

import (
	"blog-golang/db"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Erro ao carregar o arquvo .env: %v", err)
	}

	// h := &Handlers.Handler{db: db.DB}
	// mux.HandleFunc("GET /", h.PostIndex)
	// mux.HandleFunc("GET /posts/{slug}", h.PostIndex)

	db.Connect()
	http.ListenAndServe(":8000", nil)
}
