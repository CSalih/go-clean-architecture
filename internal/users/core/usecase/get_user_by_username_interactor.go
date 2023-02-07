package usecase

import (
	"github.com/CSalih/go-clean-architecture/internal/common/presenter"
)

type getUserByUsernameInteractor struct {
	gateway GetUserByUsernameGateway
}

func NewGetUserByUsernameInteractor(gateway GetUserByUsernameGateway) GetUserByUsernameUseCase {
	return &getUserByUsernameInteractor{
		gateway: gateway,
	}
}

func (i getUserByUsernameInteractor) Handle(query GetUserByUsernameQuery, presenter presenter.Presenter) error {
	user, err := i.gateway.GetByUsername(query)
	if err != nil {
		return presenter.OnError(err)
	}
	return presenter.OnSuccess(user)
}
