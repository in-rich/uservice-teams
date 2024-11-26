package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-teams/pkg/dao"
	"github.com/in-rich/uservice-teams/pkg/models"
)

type GetTeamService interface {
	Exec(ctx context.Context, in *models.GetTeamRequest) (*models.Team, error)
}

type getTeamService struct {
	getTeamRepository dao.GetTeamRepository
}

func (s *getTeamService) Exec(ctx context.Context, in *models.GetTeamRequest) (*models.Team, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(in); err != nil {
		return nil, errors.Join(ErrInvalidData, err)
	}

	teamID, err := uuid.Parse(in.TeamID)
	if err != nil {
		return nil, errors.Join(ErrInvalidData, err)
	}

	team, err := s.getTeamRepository.GetTeam(ctx, teamID)
	if err != nil {
		return nil, err
	}

	return &models.Team{
		TeamID:  team.ID.String(),
		OwnerID: team.OwnerID,
		Name:    team.Name,
	}, nil
}

func NewGetTeamService(getTeamRepository dao.GetTeamRepository) GetTeamService {
	return &getTeamService{
		getTeamRepository: getTeamRepository,
	}
}
