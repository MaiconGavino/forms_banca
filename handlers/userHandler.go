package handlers

import (
	"encoding/json"
	"forms_Banca/models"
	"net/http"
)

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

	user, err := models.Authenticate(credentials.Name, credentials.Password)
	if err != nil {
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	response := map[string]string{
		"name": user.Name,
		"role": "admin", // ajuste aqui para incluir lógica de roles
	}

	json.NewEncoder(w).Encode(response)
}
