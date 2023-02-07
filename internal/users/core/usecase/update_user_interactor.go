package usecase

import (
	"github.com/CSalih/go-clean-architecture/internal/users/core/presenter"
	"github.com/CSalih/go-clean-architecture/internal/users/core/problem"
)

type updateUserInteractor struct {
	updateUserGateway         UpdateUserGateway
	doesUsernameExistsGateway DoesUsernameExistsGateway
}

func NewUpdateUserInteractor(updateUserGateway UpdateUserGateway, doesUsernameExistsGateway DoesUsernameExistsGateway) UpdateUserUseCase {
	return &updateUserInteractor{
		updateUserGateway:         updateUserGateway,
		doesUsernameExistsGateway: doesUsernameExistsGateway,
	}
}

func (r updateUserInteractor) Handle(command UpdateUserCommand, presenter presenter.Presenter) error {
	if command.Username == "" {
		return presenter.OnError(problem.NewUsernameRequiredProblem())
	}
	if command.Status == "" {
		return presenter.OnError(problem.NewUserStatusRequiredProblem())
	}

	if exist, _ := r.doesUsernameExistsGateway.Exist(UsernameExistsQuery{command.Username}); !exist {
		return presenter.OnError(problem.NewUserNotFoundProblem())
	}

	user, err := r.updateUserGateway.Update(command)
	if err != nil {
		return presenter.OnError(err)
	}
	return presenter.OnSuccess(user)
}
