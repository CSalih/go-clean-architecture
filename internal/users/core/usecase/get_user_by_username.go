package usecase

import "github.com/CSalih/go-clean-architecture/internal/users/domain/model"

type GetUserByUsernameGateway interface {
	GetByUsername(GetUserByUsernameQuery) (model.User, error)
}

type GetUserByUsernameInteractor struct {
	gateway GetUserByUsernameGateway
}

func NewGetUserByUsernameInteractor(gateway GetUserByUsernameGateway) GetUserByUsernameInteractor {
	return GetUserByUsernameInteractor{
		gateway: gateway,
	}
}

func (i *GetUserByUsernameInteractor) Handle(query GetUserByUsernameQuery) (model.User, error) {
	user, err := i.gateway.GetByUsername(query)
	if err != nil {
		// TODO: We need to pass a presenter
		return model.User{}, err
	}
	return user, nil
}
