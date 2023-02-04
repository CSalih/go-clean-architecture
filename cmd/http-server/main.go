package main

import (
	"github.com/CSalih/go-clean-architecture/internal/users/infrastrucure/controller"
	"github.com/CSalih/go-clean-architecture/internal/users/infrastrucure/repository"
	"github.com/CSalih/go-clean-architecture/pkg/router"
	"net/http"
)

var (
	userRepository = repository.NewInMemoryUserRepository()
	userController = controller.NewUserController(userRepository)
	r              = router.NewGinRouter()
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
