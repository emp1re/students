package service

import "go.uber.org/zap"

type Service interface {
}
type Repository interface {
	*zap.Logger
	sqlc
}

func NewService() *Repository {
	return &Repository{}
}
