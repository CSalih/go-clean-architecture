package controller

import (
	"fmt"
	"github.com/CSalih/go-clean-architecture/internal/users/domain/model"
	"github.com/CSalih/go-clean-architecture/internal/users/infrastrucure/repository"
	"github.com/CSalih/go-clean-architecture/pkg/router"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	userRepository = repository.NewInMemoryUserRepository()
	controller     = NewUserController(userRepository)
	testRouter     = router.NewChiRouter()
)

func init() {
	controller.AddRoutes(testRouter)
}

func TestRoutes_getAllUsers(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users", nil)
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "[]", w.Body.String())
}

func Test_addUser(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/users/salih", nil)
	testRouter.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Contains(t, w.Body.String(), "salih")
	assert.Contains(t, w.Body.String(), "new")
}

func Test_getUserByUsername(t *testing.T) {
	_, _ = userRepository.Save(model.User{
		Username: "salih",
		Status:   "new",
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users/salih", nil)
	testRouter.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "salih")
	assert.Contains(t, w.Body.String(), "new")
}

func Test_updateStatusOfUser(t *testing.T) {
	_, _ = userRepository.Save(model.User{
		Username: "salih",
		Status:   "new",
	})

	expectedStatus := "tester"
	payload := fmt.Sprintf(`{"status":"%s"}`, expectedStatus)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/users/salih", strings.NewReader(payload))
	testRouter.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "salih")
	assert.Contains(t, w.Body.String(), expectedStatus)
}
