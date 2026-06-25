package main

import (
	"blog-golang/db"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

	var portNumber = ":8000"

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Erro ao carregar o arquvo .env: %v", err)
	}

	db.Connect()

	// h := &Handlers.Handler{db: db.DB}
	// mux.HandleFunc("GET /", h.PostIndex)
	// mux.HandleFunc("GET /posts/{slug}", h.PostIndex)

	err = http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor: ", err)
	}
}
