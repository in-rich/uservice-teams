package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-teams/pkg/dao"
	"github.com/in-rich/uservice-teams/pkg/models"
)

type SetTeamOwnerService interface {
	Exec(ctx context.Context, in *models.SetTeamOwnerRequest) (*models.Team, error)
}

type setTeamOwnerService struct {
	setTeamOwnerRepository dao.SetTeamOwnerRepository
}

func (s *setTeamOwnerService) Exec(ctx context.Context, in *models.SetTeamOwnerRequest) (*models.Team, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(in); err != nil {
		return nil, errors.Join(ErrInvalidData, err)
	}

	teamID, err := uuid.Parse(in.TeamID)
	if err != nil {
		return nil, errors.Join(ErrInvalidData, err)
	}

	team, err := s.setTeamOwnerRepository.SetTeamOwner(ctx, teamID, in.UserID)
	if err != nil {
		return nil, err
	}

	return &models.Team{
		TeamID:  team.ID.String(),
		Name:    team.Name,
		OwnerID: team.OwnerID,
	}, nil
}

func NewSetTeamOwnerService(setTeamOwnerRepository dao.SetTeamOwnerRepository) SetTeamOwnerService {
	return &setTeamOwnerService{
		setTeamOwnerRepository: setTeamOwnerRepository,
	}
}
