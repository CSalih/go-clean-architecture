package controller

import (
	"github.com/CSalih/go-clean-architecture/internal/router"
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"net/http"
)

type getUserByUsernameHandler struct {
	getUserByUsernameUseCase usecase.GetUserByUsernameUseCase
}

func (h getUserByUsernameHandler) Handle(ctx *router.Context) {
	user, err := h.getUserByUsernameUseCase.Handle(usecase.GetUserByUsernameQuery{
		Username: ctx.Params["name"],
	})
	if err != nil {
		_ = ctx.ProblemJson(err)
		return
	}

	_ = ctx.Json(http.StatusOK, user)
}
