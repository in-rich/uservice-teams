package services

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-teams/pkg/dao"
)

type DeleteTeamService interface {
	Exec(ctx context.Context, teamID string) error
}

type deleteTeamService struct {
	deleteTeamRepository dao.DeleteTeamRepository
}

func (s *deleteTeamService) Exec(ctx context.Context, rawTeamID string) error {
	teamID, err := uuid.Parse(rawTeamID)
	if err != nil {
		return errors.Join(ErrInvalidData, err)
	}

	if err := s.deleteTeamRepository.DeleteTeam(ctx, teamID); err != nil {
		return err
	}

	return nil
}

func NewDeleteTeamService(deleteTeamRepository dao.DeleteTeamRepository) DeleteTeamService {
	return &deleteTeamService{
		deleteTeamRepository: deleteTeamRepository,
	}
}
