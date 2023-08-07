package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
	"net-api.com/internal/middleware"
)

type AuthorizationMiddlewareStub struct {
	suite.Suite
}

func TestBookController(t *testing.T) {
	suite.Run(t, new(AuthorizationMiddlewareStub))
}

func (stub *AuthorizationMiddlewareStub) SetupSuite() {
}

func (stub *AuthorizationMiddlewareStub) TestAuthorization() {
	handlerCalled := false

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerCalled = true
	})

	req, _ := http.NewRequest("GET", "localhost:3001/some-endpoint", nil)
	req.Header.Set("Api-Key", "any_key")

	rr := httptest.NewRecorder()

	middleware := middleware.Authorization(nextHandler)
	middleware.ServeHTTP(rr, req)

	stub.True(handlerCalled, "next handler should have been called")
	stub.Equal(http.StatusOK, rr.Code)
}
