package controller

import (
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"github.com/CSalih/go-clean-architecture/pkg/router"
	"net/http"
)

type GetAllUsersHandler struct {
	getAllUsersUseCase usecase.GetAllUsersUseCase
}

func (h GetAllUsersHandler) Handle(ctx *router.Context) {
	user, err := h.getAllUsersUseCase.Handle(usecase.GetAllUsersQuery{})
	if err != nil {
		panic(err)
	}

	_ = ctx.Json(http.StatusOK, user)
}
