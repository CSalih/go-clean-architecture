package usecase

import (
	"github.com/CSalih/go-clean-architecture/internal/common/presenter"
	"github.com/CSalih/go-clean-architecture/internal/users/domain/model"
)

type AddNewUserGateway interface {
	AddNewUser(AddUserCommand) (model.User, error)
}

type DoesUsernameExistsGateway interface {
	Exist(UsernameExistsQuery) (bool, error)
}

type AddUserUseCase interface {
	Handle(AddUserCommand, presenter.Presenter) error
}

type AddUserCommand struct {
	Username string
	Status   string
}

type UsernameExistsQuery struct {
	Username string
}
