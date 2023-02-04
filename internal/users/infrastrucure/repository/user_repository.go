package repository

import "github.com/CSalih/go-clean-architecture/internal/users/domain/model"

type UserRepository interface {
	Save(User model.User) (model.User, error)
	Exists(Username string) bool
	FindAll() ([]model.User, error)
	FindByUsername(Username string) (model.User, error)
	Update(User model.User) (model.User, error)
}
