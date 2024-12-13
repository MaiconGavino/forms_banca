package routes

import (
	"github.com/MaiconGavino/forms_Banca/handlers"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(
		"/user",
		users.HandlerUsers,
	)
}
