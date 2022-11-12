package users

import (
	"context"
	"grpc_users_v1/internal/grpc/pb"
)

const (
	findUserByEmailQuery = `select id, name, email from users where email like $1`
)

func (s *UserServer) GetByEmail(ctx context.Context, req *pb.FindUsersByEmailRequest) (*pb.FindUsersByEmailResponse, error) {
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
