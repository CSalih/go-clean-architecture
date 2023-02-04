package usecase

import "github.com/CSalih/go-clean-architecture/internal/users/domain/model"

type GetAllUsersGateway interface {
	GetAll(GetAllUsersQuery) ([]model.User, error)
}

type GetAllUsersInteractor interface {
	Handle(GetAllUsersQuery) ([]model.User, error)
}

type getAllUsersInteractor struct {
	gateway GetAllUsersGateway
}

func NewGetAllUsersInteractor(gateway GetAllUsersGateway) GetAllUsersInteractor {
	return getAllUsersInteractor{
		gateway: gateway,
	}
}

func (r getAllUsersInteractor) Handle(query GetAllUsersQuery) ([]model.User, error) {
	users, err := r.gateway.GetAll(query)
	if err != nil {
		// TODO: We need to pass a presenter
		return []model.User{}, err
	}
	return users, nil
}
