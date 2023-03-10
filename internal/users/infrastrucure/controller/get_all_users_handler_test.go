package controller

import (
	"github.com/CSalih/go-clean-architecture/internal/common/router"
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

	err := handler.Handle(&router.Context{
		Request: req,
		Writer:  w,
		Params:  map[string]string{},
	})

	suite.NoError(err)
	suite.Equal(200, w.Code)
	suite.Equal("[]", w.Body.String())
}
