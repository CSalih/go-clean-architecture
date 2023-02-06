package controller

import (
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"github.com/CSalih/go-clean-architecture/pkg/router"
	"net/http"
	"net/http/httptest"
)

func (suite *ControllerIntegrationTestSuite) TestGetUserByUsernameHandler_Handle() {
	_, _ = userRepository.AddNewUser(usecase.AddUserCommand{
		Username: "salih",
		Status:   "new",
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users/salih", nil)

	handler := getUserByUsernameHandler{
		getUserByUsernameUseCase: getUserByUsernameUseCase,
	}

	handler.Handle(&router.Context{
		Request: req,
		Writer:  w,
		Params:  map[string]string{"name": "salih"},
	})

	suite.Equal(200, w.Code)
	suite.Contains(w.Body.String(), "salih")
	suite.Contains(w.Body.String(), "new")
}
