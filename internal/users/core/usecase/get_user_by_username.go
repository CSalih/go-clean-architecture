package usecase

import (
	"github.com/CSalih/go-clean-architecture/internal/common/presenter"
	"github.com/CSalih/go-clean-architecture/internal/users/domain/model"
)

type GetUserByUsernameGateway interface {
	GetByUsername(GetUserByUsernameQuery) (model.User, error)
}

type GetUserByUsernameUseCase interface {
	Handle(GetUserByUsernameQuery, presenter.Presenter) error
}

type GetUserByUsernameQuery struct {
	Username string
}
