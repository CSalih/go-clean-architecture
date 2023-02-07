package usecase

import (
	"github.com/CSalih/go-clean-architecture/internal/common/presenter"
)

type getAllUsersInteractor struct {
	gateway GetAllUsersGateway
}

// NewGetAllUsersInteractor creates a new instance witch implements GetAllUsersUseCase
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
