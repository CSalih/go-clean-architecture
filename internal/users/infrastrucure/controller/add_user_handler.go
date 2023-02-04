package controller

import (
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"github.com/CSalih/go-clean-architecture/pkg/router"
	"net/http"
)

type AddUserHandler struct {
	interactor usecase.AddUserInteractor
}

func (h AddUserHandler) Handle(ctx *router.Context) {
	user, err := h.interactor.Handle(usecase.AddUserCommand{
		Username: ctx.Params["name"],
		Status:   "new",
	})
	if err != nil {
		panic(err)
	}

	_ = ctx.Json(http.StatusCreated, user)
}
