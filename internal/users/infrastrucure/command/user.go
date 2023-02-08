package command

import (
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"github.com/CSalih/go-clean-architecture/internal/users/infrastrucure/repository"
	"github.com/spf13/cobra"
)

var (
	usersFilepath string
	baseUrl       string
)

func init() {
	UserCommand.PersistentFlags().StringVar(&usersFilepath, "users-file", "data/users.json", "Filepath to the users file")
	UserCommand.PersistentFlags().StringVar(&baseUrl, "url", "http://localhost:8080/api/v1", "Base URL of the API")

	/* Compose application */
	userRepository := repository.NewFileSystemUserRepository(usersFilepath)
	addUserInteractor := usecase.NewAddUserInteractor(userRepository, userRepository)
	getUserByUsernameInteractor := usecase.NewGetUserByUsernameInteractor(userRepository)
	getAllUsersInteractor := usecase.NewGetAllUsersInteractor(userRepository)
	updateUserInteractor := usecase.NewUpdateUserInteractor(userRepository, userRepository)

	/* Compose command */
	UserCommand.AddCommand(newUserAddCommand(addUserInteractor))
	userGetCommand := newUserGetCommand(getUserByUsernameInteractor)
	userGetCommand.AddCommand(newUserGetAllCommand(getAllUsersInteractor))
	UserCommand.AddCommand(userGetCommand)
	UserCommand.AddCommand(newUserUpdateCommand(updateUserInteractor))
}

var UserCommand = &cobra.Command{
	Use:   "user",
	Short: "Commands to manipulate user data",
}
