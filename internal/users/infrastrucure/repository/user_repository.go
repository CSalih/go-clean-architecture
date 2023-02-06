package repository

import (
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
)

type UserRepository interface {
	usecase.AddNewUserGateway
	usecase.DoesUsernameExistsGateway
	usecase.UpdateUserGateway
	usecase.GetAllUsersGateway
	usecase.GetUserByUsernameGateway
}
