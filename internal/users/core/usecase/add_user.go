package usecase

import (
	"github.com/CSalih/go-clean-architecture/internal/users/domain/model"
)

type SaveUserGateway interface {
	AddNewUser(AddUserCommand) (model.User, error)
}

type AddUserInteractor struct {
	gateway SaveUserGateway
}

func NewAddUserInteractor(gateway SaveUserGateway) AddUserInteractor {
	return AddUserInteractor{
		gateway: gateway,
	}
}

func (r AddUserInteractor) Handle(command AddUserCommand) (model.User, error) {
	// TODO: Add validation here

	// TODO: check if user already exists

	user, err := r.gateway.AddNewUser(command)
	if err != nil {
		// TODO: We need to pass a presenter
		return model.User{}, err
	}
	return user, nil
}
