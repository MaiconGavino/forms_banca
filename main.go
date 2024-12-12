package main

import (
	"forms_Banca/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	database.Connect()

	router := mux.NewRouter()
	routes.RegisterUserRoutes(router)
	routes.RegisterStartupRoutes(router)
	routes.RegisterEvaluetionRoutes(router)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8000", router))
}
