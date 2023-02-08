package command

import (
	"github.com/CSalih/go-clean-architecture/internal/common/presenter"
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"github.com/spf13/cobra"
)

func newUserUpdateCommand(updateUserUseCase usecase.UpdateUserUseCase) *cobra.Command {
	var userUpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update a user",
		Run: func(cmd *cobra.Command, args []string) {
			command := usecase.UpdateUserCommand{}
			if username, _ := cmd.Flags().GetString("username"); username != "" {
				command.Username = username
			}
			if status, _ := cmd.Flags().GetString("status"); status != "" {
				command.Status = status
			}

			consolePresenter := presenter.NewConsolePresenter()
			err := updateUserUseCase.Handle(command, consolePresenter)
			if err != nil {
				panic(err)
			}
		},
	}

	userUpdateCmd.Flags().String("username", "", "Username of the user")
	userUpdateCmd.Flags().String("status", "", "Status of the user")
	_ = userUpdateCmd.MarkFlagRequired("username")

	return userUpdateCmd
}
