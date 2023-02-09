package repository

import (
	"encoding/json"
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"github.com/CSalih/go-clean-architecture/internal/users/domain/model"
	"io/ioutil"
	"os"
	filepathUtil "path/filepath"
)

type fileSystemUserRepository struct {
	filepath string
}

// NewFileSystemUserRepository creates a new file system user repository instance
// It a proxy to the in memory user repository with the ability to persist data to a file
func NewFileSystemUserRepository(filepath string) UserRepository {
	return &fileSystemUserRepository{
		filepath: filepath,
	}
}

func (r *fileSystemUserRepository) AddNewUser(command usecase.AddUserCommand) (user model.User, err error) {
	repository := loadUserRepository(r.filepath)
	user, err = repository.AddNewUser(command)
	r.persistData(repository)
	return
}

func (r *fileSystemUserRepository) Update(command usecase.UpdateUserCommand) (user model.User, err error) {
	repository := loadUserRepository(r.filepath)
	user, err = repository.Update(command)
	r.persistData(repository)
	return
}

func (r *fileSystemUserRepository) GetAll(query usecase.GetAllUsersQuery) ([]model.User, error) {
	repository := loadUserRepository(r.filepath)
	return repository.GetAll(query)
}

func (r *fileSystemUserRepository) GetByUsername(query usecase.GetUserByUsernameQuery) (model.User, error) {
	repository := loadUserRepository(r.filepath)
	return repository.GetByUsername(query)
}

func (r *fileSystemUserRepository) Exist(query usecase.UsernameExistsQuery) (bool, error) {
	repository := loadUserRepository(r.filepath)
	return repository.Exist(query)
}

func loadUserRepository(filepath string) UserRepository {
	var data = make([]model.User, 0)
	jsonBytes, err := ioutil.ReadFile(filepath)
	if err == nil {
		err = json.Unmarshal(jsonBytes, &data)
		if err != nil {
			panic(err)
		}
	} else if !os.IsNotExist(err) {
		panic(err)
	}

	usernameIndex := make(map[string]model.User, len(data))
	for _, value := range data {
		usernameIndex[value.Username] = value
	}

	return &userRepository{usernameIndex}
}

func (r *fileSystemUserRepository) persistData(repository UserRepository) {
	data, err := repository.GetAll(usecase.GetAllUsersQuery{})
	if err != nil {
		panic(err)
	}
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(r.filepath); os.IsNotExist(err) {
		err := os.MkdirAll(filepathUtil.Dir(r.filepath), os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	err = ioutil.WriteFile(r.filepath, jsonBytes, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
