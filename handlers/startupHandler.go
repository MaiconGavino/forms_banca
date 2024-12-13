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
		http.Error(w, "Erro ao criar startup", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func EvaluateStartupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var evaluation models.Evaluation
	if err := json.NewDecoder(r.Body).Decode(&evaluation); err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	if err := evaluation.Create(); err != nil {
		http.Error(w, "Erro ao registrar avaliação", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
