package usecase

import "github.com/CSalih/go-clean-architecture/internal/users/domain/model"

type GetAllUsersGateway interface {
	GetAll(GetAllUsersQuery) ([]model.User, error)
}

type GetAllUsersUseCase interface {
	Handle(GetAllUsersQuery) ([]model.User, error)
}

type GetAllUsersQuery struct {
}
