package models

type UpdateTeamMemberRequest struct {
	TeamID string `json:"team_id" validate:"required,max=255"`
	UserID string `json:"user_id" validate:"required,max=255"`
	Role   string `json:"role" validate:"required,oneof=admin member"`
}
