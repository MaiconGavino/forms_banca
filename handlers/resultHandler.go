package handlers

import (
	"encoding/json"
	"forms_Banca/models"
	"net/http"
)

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
