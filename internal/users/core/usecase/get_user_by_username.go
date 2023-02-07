package usecase

import (
	"github.com/CSalih/go-clean-architecture/internal/common/presenter"
	"github.com/CSalih/go-clean-architecture/internal/users/domain/model"
)

// GetUserByUsernameGateway is the data store gateway for getting a user by username.
type GetUserByUsernameGateway interface {
	// GetByUsername returns a user by username if found.
	GetByUsername(GetUserByUsernameQuery) (model.User, error)
}

// GetUserByUsernameUseCase is the use case for getting a user by username.
type GetUserByUsernameUseCase interface {
	// Handle returns a user by username if found.
	Handle(GetUserByUsernameQuery, presenter.Presenter) error
}

// GetUserByUsernameQuery is the query for getting a user by username.
type GetUserByUsernameQuery struct {
	Username string
}
