package controller

import (
	"github.com/CSalih/go-clean-architecture/internal/common/router"
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"github.com/CSalih/go-clean-architecture/internal/users/infrastrucure/presenter"
	"net/http"
)

type getAllUsersHandler struct {
	getAllUsersUseCase usecase.GetAllUsersUseCase
}

func (h getAllUsersHandler) Handle(ctx *router.Context) error {
	jsonPresenter := presenter.NewJsonHttpPresenter(ctx.Writer, http.StatusOK)
	query := usecase.GetAllUsersQuery{}
	return h.getAllUsersUseCase.Handle(query, jsonPresenter)
}
