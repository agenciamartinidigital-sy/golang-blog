package main

import (
	"blog-golang/db"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Erro ao carregar o arquvo .env: %v", err)
	}

	db.Connect()
	log.Println("Configurado com sucesso!")
}
