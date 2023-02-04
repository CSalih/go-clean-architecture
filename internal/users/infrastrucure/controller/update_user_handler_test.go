package controller

import (
	"fmt"
	"github.com/CSalih/go-clean-architecture/internal/users/core/usecase"
	"github.com/CSalih/go-clean-architecture/pkg/router"
	"net/http"
	"net/http/httptest"
	"strings"
)

func (suite *ControllerIntegrationTestSuite) TestUpdateUserHandler_Handle() {
	_, _ = userRepository.AddNewUser(usecase.AddUserCommand{
		Username: "salih",
		Status:   "new",
	})

	expectedStatus := "tester"
	payload := fmt.Sprintf(`{"status":"%s"}`, expectedStatus)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/users/salih", strings.NewReader(payload))

	handler := UpdateUserHandler{
		interactor: updateUserInteractor,
	}

	handler.Handle(&router.Context{
		Request: req,
		Writer:  w,
		Params:  map[string]string{"name": "salih"},
	})

	suite.Equal(200, w.Code)
	suite.Contains(w.Body.String(), "salih")
	suite.Contains(w.Body.String(), expectedStatus)
}
