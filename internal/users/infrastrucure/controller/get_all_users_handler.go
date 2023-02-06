package controller

import (
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"github.com/CSalih/go-clean-architecture/pkg/router"
	"net/http"
)

type getAllUsersHandler struct {
	getAllUsersUseCase usecase.GetAllUsersUseCase
}

func (h getAllUsersHandler) Handle(ctx *router.Context) {
	user, err := h.getAllUsersUseCase.Handle(usecase.GetAllUsersQuery{})
	if err != nil {
		_ = ctx.ProblemJson(err)
		return
	}

	_ = ctx.Json(http.StatusOK, user)
}
