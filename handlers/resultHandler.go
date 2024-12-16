package handlers

import (
	"encoding/json"
	"forms_Banca/models"
	"net/http"
)

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

func GetResultsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	results, err := models.GetResults()
	if err != nil {
		http.Error(w, "Erro ao buscar resultados", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(results)
}
