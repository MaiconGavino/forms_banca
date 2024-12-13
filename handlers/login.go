package handlers

import (
	"net/http"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	name := r.PostFormValue("name")
	password := r.PostFormValue("password")

	if name == "" || password == "" {
		http.Error(w, "Missing name or password", http.StatusBadRequest)
		return
	}

	query := `SELECT id FROM users WHERE name = $1 AND password = $2`
	var userID int
	err := database.DB.QueryRow(query, name, password).Scan(&userID)
	if err != nil {
		http.Error(w, "Invalid credential", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
