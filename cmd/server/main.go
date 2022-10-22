package main

import (
	"google.golang.org/grpc"
	"grpc_users_v1/internal/config"
	"grpc_users_v1/internal/grpc/pb"
	"grpc_users_v1/pkg/repository"
	"grpc_users_v1/pkg/users"
	"log"
	"net"
)

func main() {
	dbConf, err := config.GetDbConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := repository.NewPostgresDb(dbConf)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	srv := users.NewGRPCServer(db)
	pb.RegisterUsersServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
