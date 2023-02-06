package usecase

import "github.com/CSalih/go-clean-architecture/internal/users/domain/model"

type UpdateUserGateway interface {
	Update(UpdateUserCommand) (model.User, error)
}

type UpdateUserUseCase interface {
	Handle(UpdateUserCommand) (model.User, error)
}

type UpdateUserCommand struct {
	Username string
	Status   string
}
