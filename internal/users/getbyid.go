package users

import (
	"context"
	"grpc_users_v1/internal/grpc/pb"
)

const (
	findUserByIdQuery = `select name, email from users where id=$1`
)

func (s *UserServer) GetById(ctx context.Context, req *pb.FindUserByIdRequest) (*pb.FindUserByIdResponse, error) {
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
