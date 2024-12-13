package main

import (
	"forms_Banca/config"
	"forms_Banca/routes"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()
	routes.RegisterRoutes()

	log.Println("Servidor iniciado na porta 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
