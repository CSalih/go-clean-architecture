package repository

import (
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
)

type UserRepository interface {
	usecase.SaveUserGateway
	usecase.UpdateUserGateway
	usecase.GetAllUsersGateway
	usecase.GetUserByUsernameGateway
}
