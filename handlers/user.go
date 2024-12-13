package handlers

import "net/http"

func handleUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not allowed.", http.StatusMethodNotAllowed)
		return
	}
	name := r.PostFormValue("name")
	password := r.PostFormValue("password")

	if name == "" || password == "" {
		http.Error(w, "Missing name or password.", http.StatusBadRequest)
		return
	}
	query := `INSERT INTO users ( name, password ) VALUES ($1,$2)`
	_, err := database.DB.Exec(query, name, password)
	if err != nil {
		http.Error(w, "Failed create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully."))
}
