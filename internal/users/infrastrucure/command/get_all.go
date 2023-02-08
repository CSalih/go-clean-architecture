package command

import (
	"github.com/CSalih/go-clean-architecture/internal/common/presenter"
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"github.com/spf13/cobra"
)

func newUserGetAllCommand(getAllUsersUseCase usecase.GetAllUsersUseCase) *cobra.Command {
	var userGetAllCmd = &cobra.Command{
		Use:   "all",
		Short: "Get all users",
		Run: func(cmd *cobra.Command, args []string) {
			consolePresenter := presenter.NewConsolePresenter()
			command := usecase.GetAllUsersQuery{}
			err := getAllUsersUseCase.Handle(command, consolePresenter)
			if err != nil {
				panic(err)
			}
		},
	}

	return userGetAllCmd
}
