package main

import (
	"database/sql"
	"fmt"
	teams_pb "github.com/in-rich/proto/proto-go/teams"
	"github.com/in-rich/uservice-teams/config"
	"github.com/in-rich/uservice-teams/migrations"
	"github.com/in-rich/uservice-teams/pkg/dao"
	"github.com/in-rich/uservice-teams/pkg/handlers"
	"github.com/in-rich/uservice-teams/pkg/services"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func main() {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(config.App.Postgres.DSN)))
	db := bun.NewDB(sqldb, pgdialect.New())

	defer func() {
		_ = db.Close()
		_ = sqldb.Close()
	}()

	err := db.Ping()
	for i := 0; i < 10 && err != nil; i++ {
		time.Sleep(1 * time.Second)
		err = db.Ping()
	}

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

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.App.Server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	defer func() {
		server.GracefulStop()
		_ = listener.Close()
	}()

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
