package usecase

import (
	"github.com/CSalih/go-clean-architecture/internal/users/core/problem"
	"github.com/CSalih/go-clean-architecture/internal/users/domain/model"
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

func (r updateUserInteractor) Handle(command UpdateUserCommand) (model.User, error) {
	if command.Username == "" {
		return model.User{}, problem.UsernameRequiredProblem{}
	}
	if command.Status == "" {
		return model.User{}, problem.UserStatusRequiredProblem{}
	}

	if exist, _ := r.doesUsernameExistsGateway.Exist(UsernameExistsQuery{command.Username}); !exist {
		return model.User{}, problem.UserNotFoundProblem{}
	}

	user, err := r.updateUserGateway.Update(command)
	if err != nil {
		// TODO: We need to pass a presenter
		return model.User{}, err
	}
	return user, nil
}
