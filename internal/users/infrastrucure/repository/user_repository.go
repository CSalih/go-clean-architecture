package repository

import (
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
)

// UserRepository embeds all user data store gateway interfaces.
//
// See usecase.AddNewUserGateway, usecase.DoesUsernameExistsGateway,
// usecase.UpdateUserGateway, usecase.GetAllUsersGateway, usecase.GetUserByUsernameGateway
type UserRepository interface {
	usecase.AddNewUserGateway
	usecase.DoesUsernameExistsGateway
	usecase.UpdateUserGateway
	usecase.GetAllUsersGateway
	usecase.GetUserByUsernameGateway
}
