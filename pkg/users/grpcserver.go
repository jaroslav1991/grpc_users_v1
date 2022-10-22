package users

import (
	"context"
	"database/sql"
	"grpc_users_v1/internal/grpc/pb"
)

const (
	createUserQuery      = `insert into users (name, email, password_hash) values ($1, $2, $3) returning id`
	findUserByIdQuery    = `select name, email from users where id=$1`
	findUserByEmailQuery = `select id, name, email from users where email like $1`
)

type GRPCServer struct {
	pb.UnimplementedUsersServer
	db *sql.DB
}

func NewGRPCServer(db *sql.DB) *GRPCServer {
	return &GRPCServer{db: db}
}

func (s *GRPCServer) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	var user pb.CreateUserResponse
	rows, err := s.db.Query(createUserQuery, req.Name, req.Email, req.PasswordHash)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&user.Id); err != nil {
			return nil, err
		}
	}
	createUser := &pb.CreateUserResponse{Id: user.GetId(), Name: req.GetName(), Email: req.GetEmail()}

	return createUser, nil
}

func (s *GRPCServer) GetById(ctx context.Context, req *pb.FindUserByIdRequest) (*pb.FindUserByIdResponse, error) {
	var user pb.FindUserByIdResponse
	rows, err := s.db.Query(findUserByIdQuery, req.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&user.Name, &user.Email); err != nil {
			return nil, err
		}
	}
	getUser := &pb.FindUserByIdResponse{Id: req.Id, Name: user.Name, Email: user.Email}
	return getUser, nil
}

func (s *GRPCServer) GetByEmail(ctx context.Context, req *pb.FindUsersByEmailRequest) (*pb.FindUsersByEmailResponse, error) {
	rows, err := s.db.Query(findUserByEmailQuery, req.Email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sl []*pb.ReadUser
	for rows.Next() {
		var user pb.ReadUser
		if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		sl = append(sl, &user)
	}
	getUsers := &pb.FindUsersByEmailResponse{Users: sl}
	return getUsers, nil
}
