package models

type JoinTeamWithInvitationRequest struct {
	Code   string `json:"invitation_code" validate:"required,max=255"`
	UserID string `json:"user_id" validate:"required,max=255"`
}

type JoinTeamWithInvitationResponse struct {
	TeamID string `json:"team_id"`
}
