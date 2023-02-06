package controller

import (
	"github.com/CSalih/go-clean-architecture/internal/common/router"
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"github.com/CSalih/go-clean-architecture/internal/users/infrastrucure/presenter"
	"net/http"
)

type addUserHandler struct {
	addUserUseCase usecase.AddUserUseCase
}

func (h addUserHandler) Handle(ctx *router.Context) error {
	jsonPresenter := presenter.NewJsonHttpPresenter(ctx.Writer, http.StatusCreated)
	command := usecase.AddUserCommand{
		Username: ctx.Params["name"],
		Status:   "new",
	}
	return h.addUserUseCase.Handle(command, jsonPresenter)
}
