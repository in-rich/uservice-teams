package models

type CreateTeamRequest struct {
	Name      string `json:"name" validate:"required,max=64"`
	CreatorID string `json:"creator_id" validate:"required,max=255"`
}
