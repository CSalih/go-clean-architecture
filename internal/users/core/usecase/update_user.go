package usecase

import (
	"github.com/CSalih/go-clean-architecture/internal/common/presenter"
	"github.com/CSalih/go-clean-architecture/internal/users/domain/model"
)

// UpdateUserGateway is the data store gateway for updating a user.
type UpdateUserGateway interface {
	// Update updates a user in the data store
	Update(UpdateUserCommand) (model.User, error)
}

// UpdateUserUseCase is the use case for updating a user.
type UpdateUserUseCase interface {
	// Handle updates the user data if found.
	Handle(UpdateUserCommand, presenter.Presenter) error
}

// UpdateUserCommand is the command for updating a user.
type UpdateUserCommand struct {
	Username string
	Status   string
}
