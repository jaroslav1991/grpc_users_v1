package users

import (
	"context"
	"errors"
	"grpc_users_v1/internal/grpc/pb"
	"grpc_users_v1/internal/users/validators"
)

const (
	authorizeQuery = `select id, password_hash, name from users where email=$1`
)

func (s *UserServer) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	var user pb.SignInResponse
	var passwordHash string

	if errEmail := validators.ExistEmail(s.db, req.Email); errEmail != nil {
		return &pb.SignInResponse{
			Error: &pb.Error{
				Code:    102,
				Message: "user not found",
			},
		}, nil
	}

	rows := s.db.QueryRow(authorizeQuery, req.Email)
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	if err := rows.Scan(&user.Id, &passwordHash, &user.Name); err != nil {
		return nil, err
	}

	checkPassword := CheckPasswordHash(req.PasswordHash, passwordHash)
	if !checkPassword {
		return nil, errors.New("invalid password")
	}
	token, err := generateJWT(user.Id, user.Name)
	if err != nil {
		return nil, err
	}

	//tokenString := &pb.Token{Token: token}
	user.Token = &pb.Token{Token: token}
	//res := &pb.SignInResponse{Token: tokenString, Id: user.Id}
	return &user, nil
}
