package usecase

import (
	"github.com/CSalih/go-clean-architecture/internal/users/core/problem"
	"github.com/CSalih/go-clean-architecture/internal/users/infrastrucure/presenter"
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

func (r addUserInteractor) Handle(command AddUserCommand, presenter presenter.Presenter) error {
	if exist, _ := r.doesUsernameExistsGateway.Exist(UsernameExistsQuery{command.Username}); exist {
		return presenter.OnError(problem.NewUserExistsProblem())
	}

	user, err := r.addNewUserGateway.AddNewUser(command)
	if err != nil {
		return presenter.OnError(err)
	}
	return presenter.OnSuccess(user)
}
