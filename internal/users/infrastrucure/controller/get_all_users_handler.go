package controller

import (
	"github.com/CSalih/go-clean-architecture/internal/common/presenter"
	"github.com/CSalih/go-clean-architecture/internal/common/router"
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"net/http"
)

type getAllUsersHandler struct {
	getAllUsersUseCase usecase.GetAllUsersUseCase
}

func (h getAllUsersHandler) Handle(ctx *router.Context) error {
	jsonPresenter := presenter.NewJsonResponsePresenter(ctx.Writer, http.StatusOK)
	query := usecase.GetAllUsersQuery{}
	return h.getAllUsersUseCase.Handle(query, jsonPresenter)
}
