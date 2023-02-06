package controller

import (
	"github.com/CSalih/go-clean-architecture/pkg/router"
	"net/http"
	"net/http/httptest"
)

func (suite *ControllerIntegrationTestSuite) TestGetAllUsersHandler_Handle() {
	setupController()

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users", nil)
	w := httptest.NewRecorder()

	handler := getAllUsersHandler{
		getAllUsersUseCase: getAllUsersUseCase,
	}

	handler.Handle(&router.Context{
		Request: req,
		Writer:  w,
		Params:  map[string]string{},
	})

	suite.Equal(200, w.Code)
	suite.Equal("[]", w.Body.String())
}
