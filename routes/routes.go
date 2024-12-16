package routes

import (
	"forms_Banca/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/startup/register", handlers.RegisterStartupHandler)
	http.HandleFunc("/startup/list", handlers.ListStartupsHandler)
	http.HandleFunc("/startup/evaluate", handlers.EvaluateStartupHandler)
	http.HandleFunc("/result", handlers.GetResultsHandler)
}
