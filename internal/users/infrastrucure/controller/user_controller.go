package controller

import (
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"github.com/CSalih/go-clean-architecture/pkg/router"
)

type userController struct {
	addUserHandler           AddUserHandler
	getAllUsersHandler       GetAllUsersHandler
	getUserByUsernameHandler GetUserByUsernameHandler
	updateUserHandler        UpdateUserHandler
}

type UserController interface {
	// AddRoutes adds the routes of the controller to the given router
	AddRoutes(r router.Router)
}

func NewUserController(
	addUserUseCase usecase.AddUserUseCase,
	getAllUsersUseCase usecase.GetAllUsersUseCase,
	getUserByUsernameUseCase usecase.GetUserByUsernameUseCase,
	updateUserUseCase usecase.UpdateUserUseCase,
) UserController {
	return &userController{
		addUserHandler: AddUserHandler{
			addUserUseCase: addUserUseCase,
		},
		getAllUsersHandler: GetAllUsersHandler{
			getAllUsersUseCase: getAllUsersUseCase,
		},
		getUserByUsernameHandler: GetUserByUsernameHandler{
			getUserByUsernameUseCase: getUserByUsernameUseCase,
		},
		updateUserHandler: UpdateUserHandler{
			updateUserUseCase: updateUserUseCase,
		},
	}
}

func (c *userController) AddRoutes(r router.Router) {
	r.Get("/api/v1/users", c.getAllUsersHandler.Handle)
	r.Post("/api/v1/users/{name}", c.addUserHandler.Handle)
	r.Put("/api/v1/users/{name}", c.updateUserHandler.Handle)
	r.Get("/api/v1/users/{name}", c.getUserByUsernameHandler.Handle)
}
