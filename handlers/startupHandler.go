package handlers

import (
	"encoding/json"
	"forms_Banca/models"
	"net/http"
)

func RegisterStartupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var startup models.Startup
	if err := json.NewDecoder(r.Body).Decode(&startup); err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	if err := startup.Create(); err != nil {
		http.Error(w, "Erro ao registrar startup", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func ListStartupsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	startups, err := models.GetAllStartups()
	if err != nil {
		http.Error(w, "Erro ao listar startups", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(startups)
}
