package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	database.Connect()

	// Configuração dos endpoints
	http.HandleFunc("/users", handleUsers)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/startups", handleStartups)
	http.HandleFunc("/evaluations", handleEvaluations)
	http.HandleFunc("/results", handleResults)

	// Iniciar o servidor
	port := ":8080"
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
