package service_test

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"net-api.com/internal/domain/entities"
	"net-api.com/internal/infra/grpc"
	pb "net-api.com/internal/infra/grpc/proto"
	"net-api.com/internal/service"
)

type MockBookService struct {
	mock.Mock
}

func (m *MockBookService) GetRandomBook(ctx context.Context, in *pb.GetBookRandomRequest) (pb.PrivateBookService_GetRandomBookClient, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(pb.PrivateBookService_GetRandomBookClient), args.Error(1)
}

func (m *MockBookService) GetBookDetail(ctx context.Context, in *pb.GetBookDetailsRequest) (*pb.GetBookDetailsResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*pb.GetBookDetailsResponse), args.Error(1)
}

type BookServiceStub struct {
	suite.Suite
	// client  *MockBookService
	service *service.BookService
}

func TestBookTestSuite(t *testing.T) {
	suite.Run(t, new(BookServiceStub))
}

func (s *BookServiceStub) SetupSuite() {
	conn, err := grpc.NewBookClient()
	if err != nil {
		log.Print(err.Error())
	}

	client := pb.NewPrivateBookServiceClient(conn)
	s.service = service.NewBookService(client)
}

func (srv *BookServiceStub) TestGetHash() {
	data := &entities.User{Username: "test", Password: "pass"}
	body, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "http://localhost:3001/api/v1/user/hash", bytes.NewReader(body))
	rr := httptest.NewRecorder()
	srv.service.GetHash(rr, req)
	srv.Equal(http.StatusOK, rr.Code)
}

func (srv *BookServiceStub) TestGetBookRandom() {
	req, _ := http.NewRequest("GET", "http://localhost:3001/api/v1/book/random", nil)
	req.Header.Set("Api-Key", "your_api_key")
	rr := httptest.NewRecorder()

	// mockClient := &MockBookService{}
	// defer mockClient.AssertExpectations(srv.T())

	// srv.client = mockClient
	srv.service.GetBookRandom(rr, req)

	srv.Equal(http.StatusOK, rr.Code)
}

func (srv *BookServiceStub) TestGetBookDetail() {
	req, _ := http.NewRequest("GET", "http://localhost:3001/api/v1/book/1", nil)
	req.Header.Set("Api-Key", "your_api_key")
	rr := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("bookId", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	// mockClient := &MockBookService{}
	// defer mockClient.AssertExpectations(srv.T())

	// srv.client = mockClient
	srv.service.GetBookByID(rr, req)

	srv.Equal(http.StatusOK, rr.Code)
}
