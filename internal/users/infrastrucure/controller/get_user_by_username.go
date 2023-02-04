package controller

import (
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"github.com/CSalih/go-clean-architecture/pkg/router"
	"net/http"
)

type GetUserByUsernameHandler struct {
	interactor usecase.GetUserByUsernameInteractor
}

func (h GetUserByUsernameHandler) Handle(ctx *router.Context) {
	user, err := h.interactor.Handle(usecase.GetUserByUsernameQuery{
		Username: ctx.Params["name"],
	})
	if err != nil {
		panic(err)
	}

	_ = ctx.Json(http.StatusOK, user)
}
