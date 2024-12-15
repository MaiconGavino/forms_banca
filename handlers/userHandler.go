package handlers

import (
	"encoding/json"
	"forms_Banca/config"
	"forms_Banca/models"
	"net/http"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	if err := user.Create(); err != nil {
		http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var credentials models.User
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	// Autenticar usuário
	user, err := models.Authenticate(credentials.Name, credentials.Password)
	if err != nil {
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	// Identificar se o usuário é admin
	role := "user"
	if user.Name == "admin" {
		role = "admin"
	}

	// Responder com os dados básicos do usuário
	response := map[string]string{
		"role": role,
		"name": user.Name,
	}

	json.NewEncoder(w).Encode(response)
}

func Authenticate(name, password string) (models.User, error) {
	var user models.User
	query := "SELECT id, name, password FROM users WHERE name = $1 AND password = $2"
	err := config.DB.QueryRow(query, name, password).Scan(&user.ID, &user.Name, &user.Password)
	return user, err
}
