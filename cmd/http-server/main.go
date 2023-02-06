package main

import (
	"github.com/CSalih/go-clean-architecture/internal/common/router"
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"github.com/CSalih/go-clean-architecture/internal/users/infrastrucure/controller"
	"github.com/CSalih/go-clean-architecture/internal/users/infrastrucure/repository"
	"net/http"
)

var (
	userRepository              = repository.NewInMemoryUserRepository()
	addUserInteractor           = usecase.NewAddUserInteractor(userRepository, userRepository)
	getAllUsersInteractor       = usecase.NewGetAllUsersInteractor(userRepository)
	getUserByUsernameInteractor = usecase.NewGetUserByUsernameInteractor(userRepository)
	updateUserInteractor        = usecase.NewUpdateUserInteractor(userRepository, userRepository)
	userController              = controller.NewUserController(addUserInteractor, getAllUsersInteractor, getUserByUsernameInteractor, updateUserInteractor)
	r                           = router.NewGinRouter()
)

func init() {
	userController.AddRoutes(r)
}

func main() {
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic("Could not start HTTP Server! Got: " + err.Error())
	}
}
