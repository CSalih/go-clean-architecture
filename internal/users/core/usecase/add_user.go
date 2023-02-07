package usecase

import (
	"github.com/CSalih/go-clean-architecture/internal/common/presenter"
	"github.com/CSalih/go-clean-architecture/internal/users/domain/model"
)

// AddNewUserGateway is the data store gateway to add a new user.
type AddNewUserGateway interface {
	// AddNewUser persists the user into a data store.
	AddNewUser(AddUserCommand) (model.User, error)
}

// DoesUsernameExistsGateway is the data store gateway to check if the username exists.
type DoesUsernameExistsGateway interface {
	// Exist checks if the username exists in the data store.
	Exist(UsernameExistsQuery) (bool, error)
}

// AddUserUseCase is the use case to add a new user.
type AddUserUseCase interface {
	// Handle adds a new user to the data store if the username doesn't exist.
	Handle(AddUserCommand, presenter.Presenter) error
}

// AddUserCommand is the command to add a new user.
type AddUserCommand struct {
	Username string
	Status   string
}

// UsernameExistsQuery is the query to check if the username exists.
type UsernameExistsQuery struct {
	Username string
}
