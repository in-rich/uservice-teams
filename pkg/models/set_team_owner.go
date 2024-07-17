package models

type SetTeamOwnerRequest struct {
	TeamID string `json:"team_id" validate:"required,max=255"`
	UserID string `json:"user_id" validate:"required,max=255"`
}
