package repository

import "github.com/CSalih/go-clean-architecture/internal/users/domain/model"

type userRepository struct {
	data map[string]model.User
}

func NewInMemoryUserRepository() UserRepository {
	return &userRepository{
		data: make(map[string]model.User),
	}
}

func (r userRepository) Save(user model.User) (model.User, error) {
	user.Id = len(r.data) + 1
	r.data[user.Username] = user
	return user, nil
}

func (r userRepository) Exists(username string) bool {
	_, exits := r.data[username]
	return exits
}

func (r userRepository) FindAll() ([]model.User, error) {
	v := make([]model.User, 0, len(r.data))
	for _, value := range r.data {
		v = append(v, value)
	}
	return v, nil
}

func (r userRepository) FindByUsername(username string) (model.User, error) {
	user, exits := r.data[username]
	if !exits {
		return model.User{}, &UserNotFoundError{}
	}
	return user, nil
}

func (r userRepository) Update(user model.User) (model.User, error) {
	_, exits := r.data[user.Username]
	if !exits {
		return model.User{}, &UserNotFoundError{}
	}
	r.data[user.Username] = user
	return user, nil
}
