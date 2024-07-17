package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/in-rich/uservice-teams/pkg/dao"
	"github.com/in-rich/uservice-teams/pkg/models"
)

type CreateTeamService interface {
	Exec(ctx context.Context, in *models.CreateTeamRequest) (*models.Team, error)
}

type createTeamService struct {
	createTeamRepository dao.CreateTeamRepository
}

func (s *createTeamService) Exec(ctx context.Context, in *models.CreateTeamRequest) (*models.Team, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(in); err != nil {
		return nil, errors.Join(ErrInvalidData, err)
	}

	team, err := s.createTeamRepository.CreateTeam(ctx, &dao.CreateTeamData{
		Name:    in.Name,
		OwnerID: in.CreatorID,
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

func NewCreateTeamService(createTeamRepository dao.CreateTeamRepository) CreateTeamService {
	return &createTeamService{
		createTeamRepository: createTeamRepository,
	}
}
