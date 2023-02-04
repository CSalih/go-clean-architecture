package controller

import (
	"encoding/json"
	"github.com/CSalih/go-clean-architecture/internal/users/domain/model"
	"github.com/CSalih/go-clean-architecture/internal/users/infrastrucure/repository"
	"github.com/CSalih/go-clean-architecture/pkg/router"
	"net/http"
)

type userController struct {
	repository repository.UserRepository
}

type UserController interface {
	AddRoutes(r router.Router)
}

func NewUserController(repository repository.UserRepository) UserController {
	return &userController{
		repository: repository,
	}
}

func (c *userController) AddRoutes(r router.Router) {
	r.Get("/api/v1/users", c.getAllUsers)
	r.Post("/api/v1/users/{name}", c.addUser)
	r.Put("/api/v1/users/{name}", c.updateStatusOfUser)
	r.Get("/api/v1/users/{name}", c.getUserByUsername)
}

func (c *userController) getAllUsers(ctx *router.Context) {
	users, _ := c.repository.FindAll()

	_ = ctx.Json(http.StatusOK, users)
}

func (c *userController) addUser(ctx *router.Context) {
	username := ctx.Params["name"]
	if username == "" {
		_ = ctx.Json(http.StatusBadRequest, map[string]interface{}{
			"detail": "Username is required",
			"status": http.StatusBadRequest,
		})
		return
	}

	if c.repository.Exists(username) {
		_ = ctx.Json(http.StatusBadRequest, map[string]interface{}{
			"detail": "User exists",
			"status": http.StatusBadRequest,
		})
		return
	}
	status := "new"
	user, _ := c.repository.Save(model.User{
		Username: username,
		Status:   status,
	})

	_ = ctx.Json(http.StatusCreated, user)
}

func (c *userController) getUserByUsername(ctx *router.Context) {
	username := ctx.Params["name"]
	user, err := c.repository.FindByUsername(username)
	if err != nil {
		_ = ctx.Json(http.StatusNotFound, map[string]interface{}{
			"detail": "User not found",
			"status": 404,
		})
		return
	}

	_ = ctx.Json(http.StatusOK, user)
}

func (c *userController) updateStatusOfUser(ctx *router.Context) {
	username := ctx.Params["name"]
	user, err := c.repository.FindByUsername(username)
	if err != nil {
		_ = ctx.Json(http.StatusNotFound, map[string]interface{}{
			"detail": "User not found",
			"status": 404,
		})
		return
	}

	var jsonBody struct {
		Status string `jsonBody:"status" binding:"required"`
	}

	err = json.NewDecoder(ctx.Request.Body).Decode(&jsonBody)
	if err != nil {
		_ = ctx.Json(http.StatusBadRequest, map[string]interface{}{
			"detail": err.Error(),
			"status": 404,
		})
		return
	}

	user.Status = jsonBody.Status
	user, err = c.repository.Update(user)
	if err != nil {
		_ = ctx.Json(http.StatusInternalServerError, map[string]interface{}{
			"detail": err.Error(),
			"status": 404,
		})
		return
	}

	_ = ctx.Json(http.StatusOK, user)
}
