package usecase

import "github.com/CSalih/go-clean-architecture/internal/users/domain/model"

type getUserByUsernameInteractor struct {
	gateway GetUserByUsernameGateway
}

func NewGetUserByUsernameInteractor(gateway GetUserByUsernameGateway) GetUserByUsernameUseCase {
	return &getUserByUsernameInteractor{
		gateway: gateway,
	}
}

func (i getUserByUsernameInteractor) Handle(query GetUserByUsernameQuery) (model.User, error) {
	user, err := i.gateway.GetByUsername(query)
	if err != nil {
		// TODO: We need to pass a presenter
		return model.User{}, err
	}
	return user, nil
}
