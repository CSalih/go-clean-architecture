package usecase

import (
	"github.com/CSalih/go-clean-architecture/internal/users/core/presenter"
)

type getAllUsersInteractor struct {
	gateway GetAllUsersGateway
}

func NewGetAllUsersInteractor(gateway GetAllUsersGateway) GetAllUsersUseCase {
	return &getAllUsersInteractor{
		gateway: gateway,
	}
}

func (r getAllUsersInteractor) Handle(query GetAllUsersQuery, presenter presenter.Presenter) error {
	users, err := r.gateway.GetAll(query)
	if err != nil {
		return presenter.OnError(err)
	}
	return presenter.OnSuccess(users)
}
