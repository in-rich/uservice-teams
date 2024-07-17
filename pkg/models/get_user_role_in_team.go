package models

type GetUserRoleInTeamRequest struct {
	TeamID string `json:"team_id" validate:"required"`
	UserID string `json:"user_id" validate:"required,max=255"`
}
