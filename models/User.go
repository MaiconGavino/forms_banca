package models

import "forms_Banca/config"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u *User) Create() error {
	query := "INSERT INTO users (name, password) VALUES ($1, $2)"
	_, err := config.DB.Exec(query, u.Name, u.Password)
	return err
}

func Authenticate(name, password string) (User, error) {
	var user User
	query := "SELECT id, name, password FROM users WHERE name = $1 AND password = $2"
	err := config.DB.QueryRow(query, name, password).Scan(&user.ID, &user.Name, &user.Password)
	return user, err
}
