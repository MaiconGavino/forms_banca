package models

import "forms_Banca/config"

type Evaluation struct {
	ID                        int `json:"id"`
	StartupID                 int `json:"startup_id"`
	MentorID                  int `json:"mentor_id"`
	InnovationScore           int `json:"innovation_score"`
	ImpactScore               int `json:"impact_score"`
	EconomicFeasibilityScore  int `json:"economic_feasibility_score"`
	TechnicalFeasibilityScore int `json:"technical_feasibility_score"`
	PitchScore                int `json:"pitch_score"`
}

func (e *Evaluation) Create() error {
	query := `
        INSERT INTO evaluations 
        (startup_id, mentor_id, innovation_score, impact_score, economic_feasibility_score, technical_feasibility_score, pitch_score)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `
	_, err := config.DB.Exec(query, e.StartupID, e.MentorID, e.InnovationScore, e.ImpactScore, e.EconomicFeasibilityScore, e.TechnicalFeasibilityScore, e.PitchScore)
	return err
}

func GetResults() ([]map[string]interface{}, error) {
	query := `
        SELECT 
            s.name AS startup_name,
            AVG(e.innovation_score + e.impact_score + e.economic_feasibility_score + e.technical_feasibility_score + e.pitch_score) AS average_score
        FROM evaluations e
        JOIN startups s ON e.startup_id = s.id
        GROUP BY s.id
        ORDER BY average_score ASC
    `
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		var name string
		var score float64
		if err := rows.Scan(&name, &score); err != nil {
			return nil, err
		}
		results = append(results, map[string]interface{}{
			"startup_name":  name,
			"average_score": score,
		})
	}
	return results, nil
}
