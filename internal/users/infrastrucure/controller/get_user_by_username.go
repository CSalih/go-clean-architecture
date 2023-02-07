package controller

import (
	"github.com/CSalih/go-clean-architecture/internal/common/presenter"
	"github.com/CSalih/go-clean-architecture/internal/common/router"
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"net/http"
)

type getUserByUsernameHandler struct {
	getUserByUsernameUseCase usecase.GetUserByUsernameUseCase
}

func (h getUserByUsernameHandler) Handle(ctx *router.Context) error {
	jsonPresenter := presenter.NewJsonHttpPresenter(ctx.Writer, http.StatusOK)
	query := usecase.GetUserByUsernameQuery{
		Username: ctx.Params["name"],
	}
	return h.getUserByUsernameUseCase.Handle(query, jsonPresenter)
}
