package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_users_v1/internal/grpc/pb"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
	//res, err := c.GetByEmail(context.Background(), &pb.FindUsersByEmailRequest{Email: "%gmail%"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(res)
	//
	res1, err := c.SignUp(context.Background(), &pb.SignUpRequest{Name: "admin", Email: "admin@gmail.com", PasswordHash: "admin"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res1)
	//
	//res, err := c.CreatePost(context.Background(), &pb.CreatePostRequest{Title: "test", Message: "testing message"})
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(res)
}
