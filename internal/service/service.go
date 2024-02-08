package service

type Service interface {
}
type Repository interface {
}

func NewService() Repository {
	return &Repository{}
}
