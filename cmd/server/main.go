package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_users_v1/internal/config"
	"grpc_users_v1/internal/grpc/pb"
	"grpc_users_v1/internal/handlers"
	"grpc_users_v1/internal/posts"
	"grpc_users_v1/internal/users"
	"grpc_users_v1/pkg/repository"
	"log"
	"net"
	"net/http"
)

func main() {
	go func() {
		dbConfPost, err := config.GetDbConfigPost()
		if err != nil {
			log.Fatal(err)
		}

		db, err := repository.NewPostgresDb(dbConfPost)
		if err != nil {
			log.Fatal(err)
		}
		s := grpc.NewServer()
		srv := posts.NewPostServer(db)
		pb.RegisterPostsServer(s, srv)

		l, err := net.Listen("tcp", ":8082")
		if err != nil {
			log.Fatal(err)
		}
		if err := s.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		dbConfUser, err := config.GetDbConfigUser()
		if err != nil {
			log.Fatal(err)
		}

		db, err := repository.NewPostgresDb(dbConfUser)
		if err != nil {
			log.Fatal(err)
		}
		s := grpc.NewServer()
		srv := users.NewUserServer(db)
		pb.RegisterUsersServer(s, srv)

		l, err := net.Listen("tcp", ":8081")
		if err != nil {
			log.Fatal(err)
		}
		if err := s.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()

	connUser, err := grpc.Dial(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewUsersClient(connUser)

	connPost, err := grpc.Dial(":8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	post := pb.NewPostsClient(connPost)

	http.HandleFunc("/", handlers.MainPageHandler(post))
	http.HandleFunc("/registration", handlers.Registration(client))
	http.HandleFunc("/signin", handlers.SignIn(client))
	http.HandleFunc("/logout", handlers.Logout())
	http.HandleFunc("/create-post", handlers.CreatePost(post))
	http.HandleFunc("/error-signin", handlers.ErrorSignIn())
	http.HandleFunc("/error-signup", handlers.ErrorDomain())
	http.HandleFunc("/error-create-post", handlers.ErrorCreatePost())
	http.HandleFunc("/error-password", handlers.ErrorPassword())
	http.HandleFunc("/error-email", handlers.ErrorEmail())
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatalln(err)
	}
}
