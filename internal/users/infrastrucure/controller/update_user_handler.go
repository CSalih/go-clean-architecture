package controller

import (
	"encoding/json"
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"github.com/CSalih/go-clean-architecture/pkg/router"
	"net/http"
)

type UpdateUserHandler struct {
	interactor usecase.UpdateUserInteractor
}

func (h UpdateUserHandler) Handle(ctx *router.Context) {
	var jsonBody struct {
		Status string `jsonBody:"status" binding:"required"`
	}

	err := json.NewDecoder(ctx.Request.Body).Decode(&jsonBody)
	if err != nil {
		_ = ctx.Json(http.StatusBadRequest, map[string]interface{}{
			"detail": err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	user, err := h.interactor.Handle(usecase.UpdateUserCommand{
		Username: ctx.Params["name"],
		Status:   jsonBody.Status,
	})
	if err != nil {
		panic(err)
	}

	_ = ctx.Json(http.StatusOK, user)
}