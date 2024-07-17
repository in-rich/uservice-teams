package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-teams/pkg/dao"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/in-rich/uservice-teams/pkg/models"
	"github.com/samber/lo"
)

type ListTeamMembersService interface {
	Exec(ctx context.Context, in *models.ListTeamMembersRequest) ([]*models.TeamMember, error)
}

type listTeamMembersService struct {
	listTeamMembersRepository dao.ListTeamMembersRepository
}

func (s *listTeamMembersService) Exec(ctx context.Context, in *models.ListTeamMembersRequest) ([]*models.TeamMember, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(in); err != nil {
		return nil, errors.Join(ErrInvalidData, err)
	}

	teamID, err := uuid.Parse(in.TeamID)
	if err != nil {
		return nil, errors.Join(ErrInvalidData, err)
	}

	teamMembers, err := s.listTeamMembersRepository.ListTeamMembers(ctx, teamID, in.Limit, in.Offset)
	if err != nil {
		return nil, err
	}

	return lo.Map(teamMembers, func(item *entities.TeamMember, index int) *models.TeamMember {
		return &models.TeamMember{
			TeamID: item.TeamID.String(),
			UserID: item.UserID,
			Role:   string(item.Role),
		}
	}), nil
}

func NewListTeamMembersService(listTeamMembersRepository dao.ListTeamMembersRepository) ListTeamMembersService {
	return &listTeamMembersService{
		listTeamMembersRepository: listTeamMembersRepository,
	}
}
