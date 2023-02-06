package controller

import (
	"github.com/CSalih/go-clean-architecture/internal/router"
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"net/http"
)

type addUserHandler struct {
	addUserUseCase usecase.AddUserUseCase
}

func (h addUserHandler) Handle(ctx *router.Context) {
	user, err := h.addUserUseCase.Handle(usecase.AddUserCommand{
		Username: ctx.Params["name"],
		Status:   "new",
	})
	if err != nil {
		_ = ctx.ProblemJson(err)
		return
	}

	_ = ctx.Json(http.StatusCreated, user)
}
