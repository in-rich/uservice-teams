package models

type CreateTeamMemberRequest struct {
	TeamID string `json:"teamID" validate:"required"`
	UserID string `json:"userID" validate:"required,max=255"`
	Role   string `json:"role" validate:"required,oneof=admin member"`
}
