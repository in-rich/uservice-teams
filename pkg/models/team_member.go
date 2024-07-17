package models

type TeamMember struct {
	UserID string `json:"user_id"`
	TeamID string `json:"team_id"`
	Role   string `json:"role"`
}
