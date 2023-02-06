package usecase

import (
	"github.com/CSalih/go-clean-architecture/internal/users/core/problem"
	"github.com/CSalih/go-clean-architecture/internal/users/domain/model"
)

type addUserInteractor struct {
	addNewUserGateway         AddNewUserGateway
	doesUsernameExistsGateway DoesUsernameExistsGateway
}

func NewAddUserInteractor(addNewUserGateway AddNewUserGateway, doesUsernameExistsGateway DoesUsernameExistsGateway) AddUserUseCase {
	return &addUserInteractor{
		addNewUserGateway:         addNewUserGateway,
		doesUsernameExistsGateway: doesUsernameExistsGateway,
	}
}

func (r addUserInteractor) Handle(command AddUserCommand) (model.User, error) {
	if exist, _ := r.doesUsernameExistsGateway.Exist(UsernameExistsQuery{command.Username}); exist {
		return model.User{}, problem.UserExistsProblem{}
	}

	user, err := r.addNewUserGateway.AddNewUser(command)
	if err != nil {
		// TODO: We need to pass a presenter
		return model.User{}, err
	}
	return user, nil
}
