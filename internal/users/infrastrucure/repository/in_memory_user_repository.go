package repository

import (
	"github.com/CSalih/go-clean-architecture/internal/users/core/problem"
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"github.com/CSalih/go-clean-architecture/internal/users/domain/model"
)

type userRepository struct {
	data map[string]model.User
}

func NewInMemoryUserRepository() UserRepository {
	return &userRepository{
		data: make(map[string]model.User),
	}
}

func (r userRepository) AddNewUser(command usecase.AddUserCommand) (model.User, error) {
	user := model.User{
		Id:       len(r.data) + 1,
		Username: command.Username,
		Status:   command.Status,
	}
	r.data[user.Username] = user
	return user, nil
}

func (r userRepository) Update(command usecase.UpdateUserCommand) (model.User, error) {
	_, exits := r.data[command.Username]
	if !exits {
		return model.User{}, problem.UserNotFoundProblem{}
	}
	user := model.User{
		Id:       len(r.data) + 1,
		Username: command.Username,
		Status:   command.Status,
	}
	r.data[user.Username] = user
	return user, nil
}

func (r userRepository) GetAll(_ usecase.GetAllUsersQuery) ([]model.User, error) {
	v := make([]model.User, 0, len(r.data))
	for _, value := range r.data {
		v = append(v, value)
	}
	return v, nil
}

func (r userRepository) GetByUsername(query usecase.GetUserByUsernameQuery) (model.User, error) {
	user, exits := r.data[query.Username]
	if !exits {
		return model.User{}, problem.UserNotFoundProblem{}
	}
	return user, nil
}

func (r userRepository) Exists(username string) bool {
	_, exits := r.data[username]
	return exits
}
