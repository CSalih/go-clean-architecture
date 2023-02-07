package controller

import (
	"github.com/CSalih/go-clean-architecture/internal/common/presenter"
	"github.com/CSalih/go-clean-architecture/internal/common/router"
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"net/http"
)

type addUserHandler struct {
	addUserUseCase usecase.AddUserUseCase
}

func (h addUserHandler) Handle(ctx *router.Context) error {
	jsonPresenter := presenter.NewJsonResponsePresenter(ctx.Writer, http.StatusCreated)
	command := usecase.AddUserCommand{
		Username: ctx.Params["name"],
		Status:   "new",
	}
	return h.addUserUseCase.Handle(command, jsonPresenter)
}
