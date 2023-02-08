package command

import (
	"github.com/CSalih/go-clean-architecture/internal/common/presenter"
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"github.com/spf13/cobra"
)

func newUserAddCommand(addUserUseCase usecase.AddUserUseCase) *cobra.Command {
	var userAddCmd = &cobra.Command{
		Use:   "add",
		Short: "Add a user",
		Run: func(cmd *cobra.Command, args []string) {
			username, err := cmd.Flags().GetString("username")
			if err != nil {
				panic(err)
			}
			consolePresenter := presenter.NewConsolePresenter()
			command := usecase.AddUserCommand{
				Username: username,
				Status:   "new",
			}
			err = addUserUseCase.Handle(command, consolePresenter)
			if err != nil {
				panic(err)
			}
		},
	}

	userAddCmd.Flags().String("username", "", "Username of the user")
	_ = userAddCmd.MarkFlagRequired("username")

	return userAddCmd
}
