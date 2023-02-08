package command

import (
	"github.com/CSalih/go-clean-architecture/internal/common/presenter"
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"github.com/spf13/cobra"
)

func newUserGetCommand(getUserByUsernameUseCase usecase.GetUserByUsernameUseCase) *cobra.Command {
	var userGetCmd = &cobra.Command{
		Use:   "get",
		Short: "Get a user",
		Run: func(cmd *cobra.Command, args []string) {
			username, err := cmd.Flags().GetString("username")
			if err != nil {
				panic(err)
			}
			consolePresenter := presenter.NewConsolePresenter()
			command := usecase.GetUserByUsernameQuery{
				Username: username,
			}
			err = getUserByUsernameUseCase.Handle(command, consolePresenter)
			if err != nil {
				panic(err)
			}
		},
	}

	userGetCmd.Flags().String("username", "", "Username of the user")

	return userGetCmd
}
