package config

import (
	"grpc_users_v1/pkg/repository"
)

var localDbConfUser = repository.Config{
	Host:     "localhost",
	Port:     "5432",
	Username: "postgres",
	Password: "1234",
	DBName:   "grpc_test",
	SSLMode:  "disable",
}

var localDbConfPost = repository.Config{
	Host:     "localhost",
	Port:     "5432",
	Username: "postgres",
	Password: "1234",
	DBName:   "grpc_post",
	SSLMode:  "disable",
}

func GetDbConfigUser() (repository.Config, error) {
	return localDbConfUser, nil
}

func GetDbConfigPost() (repository.Config, error) {
	return localDbConfPost, nil
}
