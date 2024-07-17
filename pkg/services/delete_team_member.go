package services

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-teams/pkg/dao"
)

type DeleteTeamMemberService interface {
	Exec(ctx context.Context, teamID string, memberID string) error
}

type deleteTeamMemberService struct {
	deleteTeamMemberRepository dao.DeleteTeamMemberRepository
}

func (s *deleteTeamMemberService) Exec(ctx context.Context, rawTeamID string, memberID string) error {
	teamID, err := uuid.Parse(rawTeamID)
	if err != nil {
		return errors.Join(ErrInvalidData, err)
	}

	if err := s.deleteTeamMemberRepository.DeleteTeamMember(ctx, teamID, memberID); err != nil {
		return err
	}

	return nil
}

func NewDeleteTeamMemberService(deleteTeamMemberRepository dao.DeleteTeamMemberRepository) DeleteTeamMemberService {
	return &deleteTeamMemberService{
		deleteTeamMemberRepository: deleteTeamMemberRepository,
	}
}
