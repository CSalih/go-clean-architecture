package controller

import (
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"github.com/CSalih/go-clean-architecture/internal/users/infrastrucure/repository"
	"github.com/CSalih/go-clean-architecture/pkg/router"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

var (
	userRepository              repository.UserRepository
	addUserInteractor           usecase.AddUserInteractor
	getAllUsersInteractor       usecase.GetAllUsersInteractor
	getUserByUsernameInteractor usecase.GetUserByUsernameInteractor
	updateUserInteractor        usecase.UpdateUserInteractor
	controller                  UserController
	testRouter                  router.Router
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
	addUserInteractor = usecase.NewAddUserInteractor(userRepository)
	getAllUsersInteractor = usecase.NewGetAllUsersInteractor(userRepository)
	getUserByUsernameInteractor = usecase.NewGetUserByUsernameInteractor(userRepository)
	updateUserInteractor = usecase.NewUpdateUserInteractor(userRepository)
	controller = NewUserController(addUserInteractor, getAllUsersInteractor, getUserByUsernameInteractor, updateUserInteractor)
	testRouter = router.NewChiRouter()

	controller.AddRoutes(testRouter)
}
