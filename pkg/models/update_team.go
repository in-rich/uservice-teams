package models

type UpdateTeamRequest struct {
	TeamID string `json:"team_id" validate:"required,max=255"`
	Name   string `json:"name" validate:"required,max=64"`
}
