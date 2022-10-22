package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_users_v1/internal/grpc/pb"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	c := pb.NewUsersClient(conn)
	//res, err := c.Create(context.Background(), &pb.CreateUserRequest{Name: "Pypsik", Email: "pypsik@gmail.com", PasswordHash: "1234"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(res)
	//res, err := c.GetById(context.Background(), &pb.FindUserByIdRequest{Id: 7})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(res)
	res, err := c.GetByEmail(context.Background(), &pb.FindUsersByEmailRequest{Email: "%gmail%"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
