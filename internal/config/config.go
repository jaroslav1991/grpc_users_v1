package config

import "grpc_users_v1/pkg/repository"

var localDbCong = repository.Config{
	Host:     "localhost",
	Port:     "5432",
	Username: "postgres",
	Password: "1234",
	DBName:   "grpc_test",
	SSLMode:  "disable",
}

func GetDbConfig() (repository.Config, error) {
	return localDbCong, nil
}
