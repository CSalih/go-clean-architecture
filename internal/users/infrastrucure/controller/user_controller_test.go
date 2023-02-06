package controller

import (
	router2 "github.com/CSalih/go-clean-architecture/internal/common/router"
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"github.com/CSalih/go-clean-architecture/internal/users/infrastrucure/repository"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

var (
	userRepository           repository.UserRepository
	addUserUseCase           usecase.AddUserUseCase
	getAllUsersUseCase       usecase.GetAllUsersUseCase
	getUserByUsernameUseCase usecase.GetUserByUsernameUseCase
	updateUserUseCase        usecase.UpdateUserUseCase
	controller               UserController
	testRouter               router2.Router
)

type ControllerIntegrationTestSuite struct {
	suite.Suite
}

func (suite *ControllerIntegrationTestSuite) SetupTest() {
	setupController()
}

func TestControllerIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(ControllerIntegrationTestSuite))
}

func setupController() {
	log.Println("Setting up controller")
	userRepository = repository.NewInMemoryUserRepository()
	addUserUseCase = usecase.NewAddUserInteractor(userRepository, userRepository)
	getAllUsersUseCase = usecase.NewGetAllUsersInteractor(userRepository)
	getUserByUsernameUseCase = usecase.NewGetUserByUsernameInteractor(userRepository)
	updateUserUseCase = usecase.NewUpdateUserInteractor(userRepository, userRepository)
	controller = NewUserController(addUserUseCase, getAllUsersUseCase, getUserByUsernameUseCase, updateUserUseCase)
	testRouter = router2.NewChiRouter()

	controller.AddRoutes(testRouter)
}
