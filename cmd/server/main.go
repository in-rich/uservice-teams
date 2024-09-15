package main

import (
	"fmt"
	"github.com/in-rich/lib-go/deploy"
	teams_pb "github.com/in-rich/proto/proto-go/teams"
	"github.com/in-rich/uservice-teams/config"
	"github.com/in-rich/uservice-teams/migrations"
	"github.com/in-rich/uservice-teams/pkg/dao"
	"github.com/in-rich/uservice-teams/pkg/handlers"
	"github.com/in-rich/uservice-teams/pkg/services"
	"log"
)

func main() {
	db, closeDB := deploy.OpenDB(config.App.Postgres.DSN)
	defer closeDB()

	if err := migrations.Migrate(db); err != nil {
		log.Fatalf("failed to migrate: %v", err)
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

	createTeamService := services.NewCreateTeamService(createTeamDAO)
	createTeamMemberService := services.NewCreateTeamMemberService(createTeamMemberDAO, getTeamDAO)
	deleteTeamService := services.NewDeleteTeamService(deleteTeamDAO)
	deleteTeamMemberService := services.NewDeleteTeamMemberService(deleteTeamMemberDAO)
	listTeamMembersService := services.NewListTeamMembersService(listTeamMembersDAO)
	getUserRoleInTeamService := services.NewGetUserRoleInTeamService(getUserRoleInTeamDAO)
	listUserTeamsService := services.NewListUserTeamsService(listUserTeamsDAO)
	setTeamOwnerService := services.NewSetTeamOwnerService(setTeamOwnerDAO)
	updateTeamService := services.NewUpdateTeamService(updateTeamDAO)
	updateTeamMemberService := services.NewUpdateTeamMemberService(updateTeamMemberDAO)

	createTeamHandler := handlers.NewCreateTeamHandler(createTeamService)
	createTeamMemberHandler := handlers.NewCreateTeamMemberHandler(createTeamMemberService)
	deleteTeamHandler := handlers.NewDeleteTeamHandler(deleteTeamService)
	deleteTeamMemberHandler := handlers.NewDeleteTeamMemberHandler(deleteTeamMemberService)
	listTeamMembersHandler := handlers.NewListTeamMembersHandler(listTeamMembersService)
	getUserRoleInTeamHandler := handlers.NewGetUserRoleInTeamHandler(getUserRoleInTeamService)
	listUserTeamsHandler := handlers.NewListUserTeamsHandler(listUserTeamsService)
	setTeamOwnerHandler := handlers.NewSetTeamOwnerHandler(setTeamOwnerService)
	updateTeamHandler := handlers.NewUpdateTeamHandler(updateTeamService)
	updateTeamMemberHandler := handlers.NewUpdateTeamMemberHandler(updateTeamMemberService)

	listener, server := deploy.StartGRPCServer(fmt.Sprintf(":%d", config.App.Server.Port), "teams")
	defer deploy.CloseGRPCServer(listener, server)

	teams_pb.RegisterCreateTeamServer(server, createTeamHandler)
	teams_pb.RegisterCreateTeamMemberServer(server, createTeamMemberHandler)
	teams_pb.RegisterDeleteTeamServer(server, deleteTeamHandler)
	teams_pb.RegisterDeleteTeamMemberServer(server, deleteTeamMemberHandler)
	teams_pb.RegisterListTeamMembersServer(server, listTeamMembersHandler)
	teams_pb.RegisterGetUserRoleInTeamServer(server, getUserRoleInTeamHandler)
	teams_pb.RegisterListUserTeamsServer(server, listUserTeamsHandler)
	teams_pb.RegisterSetTeamOwnerServer(server, setTeamOwnerHandler)
	teams_pb.RegisterUpdateTeamServer(server, updateTeamHandler)
	teams_pb.RegisterUpdateTeamMemberServer(server, updateTeamMemberHandler)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
