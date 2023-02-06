package usecase

import (
	"github.com/CSalih/go-clean-architecture/internal/users/core/problem"
	"github.com/CSalih/go-clean-architecture/internal/users/domain/model"
)

type addUserInteractor struct {
	gateway AddNewUserGateway
}

func NewAddUserInteractor(gateway AddNewUserGateway) AddUserUseCase {
	return &addUserInteractor{
		gateway: gateway,
	}
}

func (r addUserInteractor) Handle(command AddUserCommand) (model.User, error) {
	if exist, _ := r.gateway.Exist(UsernameExistsQuery{command.Username}); exist {
		return model.User{}, problem.UserExistsProblem{}
	}

	user, err := r.gateway.AddNewUser(command)
	if err != nil {
		// TODO: We need to pass a presenter
		return model.User{}, err
	}
	return user, nil
}
