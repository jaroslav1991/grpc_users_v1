package posts

import (
	"database/sql"
	"grpc_users_v1/internal/grpc/pb"
)

type PostServer struct {
	pb.UnimplementedPostsServer
	db *sql.DB
}

func NewPostServer(db *sql.DB) *PostServer {
	return &PostServer{db: db}
}
