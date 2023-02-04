package usecase

import "github.com/CSalih/go-clean-architecture/internal/users/domain/model"

type UpdateUserGateway interface {
	Update(UpdateUserCommand) (model.User, error)
}

type UpdateUserInteractor struct {
	gateway UpdateUserGateway
}

func NewUpdateUserInteractor(gateway UpdateUserGateway) UpdateUserInteractor {
	return UpdateUserInteractor{
		gateway: gateway,
	}
}

func (r UpdateUserInteractor) Handle(command UpdateUserCommand) (model.User, error) {
	// TODO: Add validation here

	// TODO: check if user already exists

	user, err := r.gateway.Update(command)
	if err != nil {
		// TODO: We need to pass a presenter
		return model.User{}, err
	}
	return user, nil
}
