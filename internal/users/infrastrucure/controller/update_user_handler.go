package controller

import (
	"encoding/json"
	"github.com/CSalih/go-clean-architecture/internal/common/router"
	"github.com/CSalih/go-clean-architecture/internal/users/core/problem"
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"github.com/CSalih/go-clean-architecture/internal/users/infrastrucure/presenter"
	"net/http"
)

type updateUserHandler struct {
	updateUserUseCase usecase.UpdateUserUseCase
}

func (h updateUserHandler) Handle(ctx *router.Context) error {
	var jsonBody struct {
		Status string `jsonBody:"status" binding:"required"`
	}

	err := json.NewDecoder(ctx.Request.Body).Decode(&jsonBody)
	if err != nil {
		return problem.Problem{
			Type:   "https://example.com/probs/invalid-json",
			Title:  "Invalid JSON payload received",
			Status: http.StatusBadRequest,
		}
	}

	jsonPresenter := presenter.NewJsonHttpPresenter(ctx.Writer, http.StatusOK)
	command := usecase.UpdateUserCommand{
		Username: ctx.Params["name"],
		Status:   jsonBody.Status,
	}
	return h.updateUserUseCase.Handle(command, jsonPresenter)
}
