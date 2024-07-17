package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-teams/pkg/dao"
	"github.com/in-rich/uservice-teams/pkg/models"
)

type UpdateTeamService interface {
	Exec(ctx context.Context, in *models.UpdateTeamRequest) (*models.Team, error)
}

type updateTeamService struct {
	updateTeamRepository dao.UpdateTeamRepository
}

func (s *updateTeamService) Exec(ctx context.Context, in *models.UpdateTeamRequest) (*models.Team, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(in); err != nil {
		return nil, errors.Join(ErrInvalidData, err)
	}

	teamID, err := uuid.Parse(in.TeamID)
	if err != nil {
		return nil, errors.Join(ErrInvalidData, err)
	}

	team, err := s.updateTeamRepository.UpdateTeam(ctx, teamID, &dao.UpdateTeamData{
		Name: in.Name,
	})
	if err != nil {
		return nil, err
	}

	return &models.Team{
		TeamID:  team.ID.String(),
		OwnerID: team.OwnerID,
		Name:    team.Name,
	}, nil
}

func NewUpdateTeamService(updateTeamRepository dao.UpdateTeamRepository) UpdateTeamService {
	return &updateTeamService{
		updateTeamRepository: updateTeamRepository,
	}
}
