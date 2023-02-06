package usecase

import "github.com/CSalih/go-clean-architecture/internal/users/domain/model"

type getAllUsersInteractor struct {
	gateway GetAllUsersGateway
}

func NewGetAllUsersInteractor(gateway GetAllUsersGateway) GetAllUsersUseCase {
	return &getAllUsersInteractor{
		gateway: gateway,
	}
}

func (r getAllUsersInteractor) Handle(query GetAllUsersQuery) ([]model.User, error) {
	users, err := r.gateway.GetAll(query)
	if err != nil {
		// TODO: We need to pass a presenter
		return []model.User{}, err
	}
	return users, nil
}
