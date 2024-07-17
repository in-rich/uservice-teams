package models

type Team struct {
	TeamID  string `json:"team_id"`
	OwnerID string `json:"owner_id"`
	Name    string `json:"name"`
}
