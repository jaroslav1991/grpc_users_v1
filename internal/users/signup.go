package users

import (
	"context"
	"grpc_users_v1/internal/grpc/pb"
	"grpc_users_v1/internal/users/validators"
	"log"
)

const (
	createUserQuery = `insert into users (name, email, password_hash) values ($1, $2, $3) returning id`
)

func (s *UserServer) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	var user pb.SignUpResponse

	if errPassword := validators.ValidatePassword(req.PasswordHash); errPassword != true {

		log.Println(errPassword)
		return &pb.SignUpResponse{
			Error: &pb.Error{
				Code:    101,
				Message: "password too short",
			},
		}, nil
	}

	if errEmail := validators.ExistEmail(s.db, req.Email); errEmail == nil {
		return &pb.SignUpResponse{
			Error: &pb.Error{
				Code:    102,
				Message: "email already exist",
			},
		}, nil
	}

	if errDomain := validators.ValidateDomain(req.Email); errDomain != true {
		return &pb.SignUpResponse{
			Error: &pb.Error{
				Code:    103,
				Message: "invalid domain",
			},
		}, nil
	}

	if errCountSymbols := validators.ValidateCountSymbol(req.Email); errCountSymbols != true {
		return &pb.SignUpResponse{
			Error: &pb.Error{
				Code:    100,
				Message: "invalid email",
			},
		}, nil
	}

	passwordHash, err := HashPassword(req.PasswordHash)
	if err != nil {
		return nil, err
	}
	req.PasswordHash = passwordHash

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

	createUser := &pb.SignUpResponse{Id: user.GetId()}
	return createUser, nil
}
