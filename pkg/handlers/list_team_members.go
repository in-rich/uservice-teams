package handlers

import (
	"context"
	"errors"
	teams_pb "github.com/in-rich/proto/proto-go/teams"
	"github.com/in-rich/uservice-teams/pkg/models"
	"github.com/in-rich/uservice-teams/pkg/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ListTeamMembersHandler struct {
	teams_pb.ListTeamMembersServer
	service services.ListTeamMembersService
}

func (h *ListTeamMembersHandler) ListTeamMembers(ctx context.Context, in *teams_pb.ListTeamMembersRequest) (*teams_pb.ListTeamMembersResponse, error) {
	teamMembers, err := h.service.Exec(ctx, &models.ListTeamMembersRequest{
		TeamID: in.GetTeamId(),
		Limit:  int(in.GetLimit()),
		Offset: int(in.GetOffset()),
	})
	if err != nil {
		if errors.Is(err, services.ErrInvalidData) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid selector: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "failed to get team members: %v", err)
	}

	members := make([]*teams_pb.TeamMember, 0, len(teamMembers))
	for _, member := range teamMembers {
		members = append(members, &teams_pb.TeamMember{
			TeamId: member.TeamID,
			UserId: member.UserID,
			Role:   member.Role,
		})
	}

	return &teams_pb.ListTeamMembersResponse{
		Members: members,
	}, nil
}

func NewListTeamMembersHandler(service services.ListTeamMembersService) *ListTeamMembersHandler {
	return &ListTeamMembersHandler{
		service: service,
	}
}
