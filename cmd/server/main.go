package main

import (
	"fmt"
	"github.com/in-rich/lib-go/deploy"
	"github.com/in-rich/lib-go/monitor"
	teams_pb "github.com/in-rich/proto/proto-go/teams"
	"github.com/in-rich/uservice-teams/config"
	"github.com/in-rich/uservice-teams/migrations"
	"github.com/in-rich/uservice-teams/pkg/dao"
	"github.com/in-rich/uservice-teams/pkg/handlers"
	"github.com/in-rich/uservice-teams/pkg/services"
	"github.com/rs/zerolog"
	"os"
)

func getLogger() monitor.GRPCLogger {
	if deploy.IsReleaseEnv() {
		return monitor.NewGCPGRPCLogger(zerolog.New(os.Stdout), "uservice-teams")
	}

	return monitor.NewConsoleGRPCLogger()
}

func main() {
	logger := getLogger()

	logger.Info("Starting server")
	db, closeDB, err := deploy.OpenDB(config.App.Postgres.DSN)
	if err != nil {
		logger.Fatal(err, "failed to connect to database")
	}
	defer closeDB()

	logger.Info("Running migrations")
	if err := migrations.Migrate(db); err != nil {
		logger.Fatal(err, "failed to migrate")
	}

	depCheck := deploy.DepsCheck{
		Dependencies: func() map[string]error {
			return map[string]error{
				"Postgres": db.Ping(),
			}
		},
		Services: deploy.DepCheckServices{
			"CreateTeam":             {"Postgres"},
			"CreateTeamMember":       {"Postgres"},
			"DeleteTeam":             {"Postgres"},
			"DeleteTeamMember":       {"Postgres"},
			"GetTeam":                {"Postgres"},
			"ListTeamMembers":        {"Postgres"},
			"GetUserRoleInTeam":      {"Postgres"},
			"ListUserTeams":          {"Postgres"},
			"SetTeamOwner":           {"Postgres"},
			"UpdateTeam":             {"Postgres"},
			"UpdateTeamMember":       {"Postgres"},
			"CreateInvitationCode":   {"Postgres"},
			"JoinTeamWithInvitation": {"Postgres"},
		},
	}

	createTeamDAO := dao.NewCreateTeamRepository(db)
	createTeamMemberDAO := dao.NewCreateTeamMemberRepository(db)
	deleteTeamDAO := dao.NewDeleteTeamRepository(db)
	deleteTeamMemberDAO := dao.NewDeleteTeamMemberRepository(db)
	getTeamDAO := dao.NewGetTeamRepository(db)
	listTeamMembersDAO := dao.NewListTeamMembersRepository(db)
	getUserRoleInTeamDAO := dao.NewGetUserRoleRepository(db)
	listUserTeamsDAO := dao.NewListUserTeamsRepository(db)
	setTeamOwnerDAO := dao.NewSetTeamOwnerRepository(db)
	updateTeamDAO := dao.NewUpdateTeamRepository(db)
	updateTeamMemberDAO := dao.NewUpdateTeamMemberRepository(db)
	createInvitationCodeDAO := dao.NewCreateInvitationCodeRepository(db)
	consumeInvitationCodeDAO := dao.NewConsumeInvitationCodeRepository(db)

	createTeamService := services.NewCreateTeamService(createTeamDAO)
	createTeamMemberService := services.NewCreateTeamMemberService(createTeamMemberDAO, getTeamDAO)
	deleteTeamService := services.NewDeleteTeamService(deleteTeamDAO)
	deleteTeamMemberService := services.NewDeleteTeamMemberService(deleteTeamMemberDAO)
	getTeamService := services.NewGetTeamService(getTeamDAO)
	listTeamMembersService := services.NewListTeamMembersService(listTeamMembersDAO)
	getUserRoleInTeamService := services.NewGetUserRoleInTeamService(getUserRoleInTeamDAO)
	listUserTeamsService := services.NewListUserTeamsService(listUserTeamsDAO)
	setTeamOwnerService := services.NewSetTeamOwnerService(setTeamOwnerDAO)
	updateTeamService := services.NewUpdateTeamService(updateTeamDAO)
	updateTeamMemberService := services.NewUpdateTeamMemberService(updateTeamMemberDAO)
	createInvitationCodeService := services.NewCreateInvitationCodeService(createInvitationCodeDAO)
	joinTeamWithInvitationService := services.NewJoinTeamWithInvitationService(consumeInvitationCodeDAO, createTeamMemberDAO)

	createTeamHandler := handlers.NewCreateTeamHandler(createTeamService, logger)
	createTeamMemberHandler := handlers.NewCreateTeamMemberHandler(createTeamMemberService, logger)
	deleteTeamHandler := handlers.NewDeleteTeamHandler(deleteTeamService, logger)
	deleteTeamMemberHandler := handlers.NewDeleteTeamMemberHandler(deleteTeamMemberService, logger)
	getTeamHandler := handlers.NewGetTeamHandler(getTeamService, logger)
	listTeamMembersHandler := handlers.NewListTeamMembersHandler(listTeamMembersService, logger)
	getUserRoleInTeamHandler := handlers.NewGetUserRoleInTeamHandler(getUserRoleInTeamService, logger)
	listUserTeamsHandler := handlers.NewListUserTeamsHandler(listUserTeamsService, logger)
	setTeamOwnerHandler := handlers.NewSetTeamOwnerHandler(setTeamOwnerService, logger)
	updateTeamHandler := handlers.NewUpdateTeamHandler(updateTeamService, logger)
	updateTeamMemberHandler := handlers.NewUpdateTeamMemberHandler(updateTeamMemberService, logger)
	createInvitationCodeHandler := handlers.NewCreateInvitationCodeHandler(createInvitationCodeService, logger)
	joinTeamWithInvitationHandler := handlers.NewJoinTeamWithInvitationHandler(joinTeamWithInvitationService, logger)

	logger.Info(fmt.Sprintf("Starting to listen on port %v", config.App.Server.Port))
	listener, server, health := deploy.StartGRPCServer(logger, config.App.Server.Port, depCheck)
	defer deploy.CloseGRPCServer(listener, server)
	go health()

	teams_pb.RegisterCreateTeamServer(server, createTeamHandler)
	teams_pb.RegisterCreateTeamMemberServer(server, createTeamMemberHandler)
	teams_pb.RegisterDeleteTeamServer(server, deleteTeamHandler)
	teams_pb.RegisterDeleteTeamMemberServer(server, deleteTeamMemberHandler)
	teams_pb.RegisterGetTeamServer(server, getTeamHandler)
	teams_pb.RegisterListTeamMembersServer(server, listTeamMembersHandler)
	teams_pb.RegisterGetUserRoleInTeamServer(server, getUserRoleInTeamHandler)
	teams_pb.RegisterListUserTeamsServer(server, listUserTeamsHandler)
	teams_pb.RegisterSetTeamOwnerServer(server, setTeamOwnerHandler)
	teams_pb.RegisterUpdateTeamServer(server, updateTeamHandler)
	teams_pb.RegisterUpdateTeamMemberServer(server, updateTeamMemberHandler)
	teams_pb.RegisterCreateInvitationCodeServer(server, createInvitationCodeHandler)
	teams_pb.RegisterJoinTeamWIthInvitationServer(server, joinTeamWithInvitationHandler)

	logger.Info("Server started")
	if err := server.Serve(listener); err != nil {
		logger.Fatal(err, "failed to serve")
	}
}
