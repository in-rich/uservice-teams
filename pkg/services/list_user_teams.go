package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/in-rich/uservice-teams/pkg/dao"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/in-rich/uservice-teams/pkg/models"
	"github.com/samber/lo"
)

type ListUserTeamsService interface {
	Exec(ctx context.Context, in *models.ListUserTeamsRequest) ([]*models.Team, error)
}

type listUserTeamsService struct {
	listUserTeamsRepository dao.ListUserTeamsRepository
}

func (s *listUserTeamsService) Exec(ctx context.Context, in *models.ListUserTeamsRequest) ([]*models.Team, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(in); err != nil {
		return nil, errors.Join(ErrInvalidData, err)
	}

	teams, err := s.listUserTeamsRepository.ListUserTeams(ctx, in.UserID, in.Limit, in.Offset)
	if err != nil {
		return nil, err
	}

	return lo.Map(teams, func(item *entities.Team, index int) *models.Team {
		return &models.Team{
			TeamID:  item.ID.String(),
			OwnerID: item.OwnerID,
			Name:    item.Name,
		}
	}), nil
}

func NewListUserTeamsService(listUserTeamsRepository dao.ListUserTeamsRepository) ListUserTeamsService {
	return &listUserTeamsService{
		listUserTeamsRepository: listUserTeamsRepository,
	}
}
