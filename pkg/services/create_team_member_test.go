package services_test

import (
	"context"
	"github.com/google/uuid"
	daomocks "github.com/in-rich/uservice-teams/pkg/dao/mocks"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/in-rich/uservice-teams/pkg/models"
	"github.com/in-rich/uservice-teams/pkg/services"
	"github.com/samber/lo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateTeamMember(t *testing.T) {
	testData := []struct {
		name string

		in *models.CreateTeamMemberRequest

		shouldCallGetTeamRepository bool
		getTeamRepositoryErr        error

		shouldCallCreateTeamMemberRepository bool
		createTeamMemberResponse             *entities.TeamMember
		createTeamMemberRepositoryErr        error

		expect    *models.TeamMember
		expectErr error
	}{
		{
			name: "CreateTeamMember",
			in: &models.CreateTeamMemberRequest{
				TeamID: "00000000-0000-0000-0000-000000000001",
				UserID: "user-id-1",
				Role:   "member",
			},
			shouldCallGetTeamRepository:          true,
			shouldCallCreateTeamMemberRepository: true,
			createTeamMemberResponse: &entities.TeamMember{
				ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				TeamID:    uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				UserID:    "user-id-1",
				Role:      "member",
				CreatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
			expect: &models.TeamMember{
				TeamID: "00000000-0000-0000-0000-000000000001",
				UserID: "user-id-1",
				Role:   "member",
			},
		},
		{
			name: "CreateTeamMemberError",
			in: &models.CreateTeamMemberRequest{
				TeamID: "00000000-0000-0000-0000-000000000001",
				UserID: "user-id-1",
				Role:   "member",
			},
			shouldCallGetTeamRepository:          true,
			shouldCallCreateTeamMemberRepository: true,
			createTeamMemberRepositoryErr:        FooErr,
			expectErr:                            FooErr,
		},
		{
			name: "GetTeamError",
			in: &models.CreateTeamMemberRequest{
				TeamID: "00000000-0000-0000-0000-000000000001",
				UserID: "user-id-1",
				Role:   "member",
			},
			shouldCallGetTeamRepository: true,
			getTeamRepositoryErr:        FooErr,
			expectErr:                   FooErr,
		},
		{
			name: "InvalidData",
			in: &models.CreateTeamMemberRequest{
				TeamID: "00000000-0000-0000-0000-000000000001",
				UserID: "user-id-1",
				Role:   "",
			},
			expectErr: services.ErrInvalidData,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			createTeamMemberRepository := daomocks.NewMockCreateTeamMemberRepository(t)
			getTeamRepository := daomocks.NewMockGetTeamRepository(t)

			if tt.shouldCallGetTeamRepository {
				getTeamRepository.On("GetTeam", context.TODO(), mock.Anything).Return(nil, tt.getTeamRepositoryErr)
			}

			if tt.shouldCallCreateTeamMemberRepository {
				createTeamMemberRepository.
					On("CreateTeamMember", context.TODO(), mock.Anything, mock.Anything, mock.Anything).
					Return(tt.createTeamMemberResponse, tt.createTeamMemberRepositoryErr)
			}

			service := services.NewCreateTeamMemberService(createTeamMemberRepository, getTeamRepository)

			res, err := service.Exec(context.TODO(), tt.in)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, res)

			createTeamMemberRepository.AssertExpectations(t)
			getTeamRepository.AssertExpectations(t)
		})
	}
}
