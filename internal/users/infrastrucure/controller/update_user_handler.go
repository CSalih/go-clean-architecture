package controller

import (
	"encoding/json"
	"github.com/CSalih/go-clean-architecture/internal/common/router"
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"net/http"
)

type updateUserHandler struct {
	updateUserUseCase usecase.UpdateUserUseCase
}

func (h updateUserHandler) Handle(ctx *router.Context) {
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

	user, err := h.updateUserUseCase.Handle(usecase.UpdateUserCommand{
		Username: ctx.Params["name"],
		Status:   jsonBody.Status,
	})
	if err != nil {
		_ = ctx.ProblemJson(err)
		return
	}

	_ = ctx.Json(http.StatusOK, user)
}
