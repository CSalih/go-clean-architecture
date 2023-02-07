package usecase

import (
	"github.com/CSalih/go-clean-architecture/internal/common/presenter"
	"github.com/CSalih/go-clean-architecture/internal/users/domain/model"
)

// GetAllUsersGateway is the data store gateway to get all users.
type GetAllUsersGateway interface {
	// GetAll returns all users from the data store.
	GetAll(GetAllUsersQuery) ([]model.User, error)
}

// GetAllUsersUseCase is the use case to get all users from the data store.
type GetAllUsersUseCase interface {
	// Handle returns all users from the data store.
	Handle(GetAllUsersQuery, presenter.Presenter) error
}

// GetAllUsersQuery is the query to get all users from the data store.
type GetAllUsersQuery struct {
}
