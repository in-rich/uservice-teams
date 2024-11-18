package models

type CreateInvitationCodeRequest struct {
	TeamID string `json:"team_id" validate:"required,max=255"`
}

type CreateInvitationCodeResponse struct {
	Code string `json:"code"`
}
