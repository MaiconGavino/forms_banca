package models

import "forms_Banca/config"

type Startup struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (s *Startup) Create() error {
	query := "INSERT INTO startups (name) VALUES ($1)"
	_, err := config.DB.Exec(query, s.Name)
	return err
}

func GetAllStartups() ([]Startup, error) {
	query := "SELECT id, name FROM startups"
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var startups []Startup
	for rows.Next() {
		var startup Startup
		if err := rows.Scan(&startup.ID, &startup.Name); err != nil {
			return nil, err
		}
		startups = append(startups, startup)
	}
	return startups, nil
}
