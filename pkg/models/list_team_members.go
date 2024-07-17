package models

type ListTeamMembersRequest struct {
	TeamID string `json:"team_id" validate:"required"`
	Limit  int    `json:"limit" validate:"required,min=1,max=1000"`
	Offset int    `json:"offset" validate:"min=0"`
}
