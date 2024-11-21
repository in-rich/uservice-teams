package models

type GetTeamRequest struct {
	TeamID string `json:"user_id" validate:"required,max=255"`
}
