package controller

import (
	"github.com/CSalih/go-clean-architecture/internal/common/router"
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
)

type userController struct {
	addUserHandler           addUserHandler
	getAllUsersHandler       getAllUsersHandler
	getUserByUsernameHandler getUserByUsernameHandler
	updateUserHandler        updateUserHandler
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
		addUserHandler: addUserHandler{
			addUserUseCase: addUserUseCase,
		},
		getAllUsersHandler: getAllUsersHandler{
			getAllUsersUseCase: getAllUsersUseCase,
		},
		getUserByUsernameHandler: getUserByUsernameHandler{
			getUserByUsernameUseCase: getUserByUsernameUseCase,
		},
		updateUserHandler: updateUserHandler{
			updateUserUseCase: updateUserUseCase,
		},
	}
}

func (c *userController) AddRoutes(r router.Router) {
	r.Get("/api/v1/users", withProblemJsonHandler(c.getAllUsersHandler.Handle))
	r.Post("/api/v1/users/{name}", withProblemJsonHandler(c.addUserHandler.Handle))
	r.Put("/api/v1/users/{name}", withProblemJsonHandler(c.updateUserHandler.Handle))
	r.Get("/api/v1/users/{name}", withProblemJsonHandler(c.getUserByUsernameHandler.Handle))
}

func withProblemJsonHandler(handler func(ctx *router.Context) error) router.HandlerFunc {
	return func(ctx *router.Context) {
		err := handler(ctx)
		if err == nil {
			return
		}
		err = ctx.ProblemJson(err)
		if err != nil {
			panic(err)
		}
	}
}
