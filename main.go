package main

import (
	"log"
	"net/http"
)

func main() {
	database.Connect()

	mux := http.NewServeMux()

	//fs := http.FileServer(http.Dir("./frontend/card"))
	//mux.Handle("/", fs)

	routes.RegisterRoutes(mux)
	http.ListenAndServe(":8080", mux)
	log.Println("Listening on port 8080")
}
