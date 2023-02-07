package usecase

import (
	"github.com/CSalih/go-clean-architecture/internal/common/presenter"
	"github.com/CSalih/go-clean-architecture/internal/users/core/problem"
)

type addUserInteractor struct {
	addNewUserGateway         AddNewUserGateway
	doesUsernameExistsGateway DoesUsernameExistsGateway
}

// NewAddUserInteractor creates a new instance witch implements AddUserUseCase
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
