package controller

import (
	"github.com/CSalih/go-clean-architecture/pkg/router"
	"net/http"
	"net/http/httptest"
)

func (suite *ControllerIntegrationTestSuite) TestAddUserHandler_Handle() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/users/salih", nil)

	addUserHandler := AddUserHandler{
		interactor: addUserInteractor,
	}

	addUserHandler.Handle(&router.Context{
		Request: req,
		Writer:  w,
		Params:  map[string]string{"name": "salih"},
	})

	suite.Equal(201, w.Code)
	suite.Contains(w.Body.String(), "salih")
	suite.Contains(w.Body.String(), "new")
}
