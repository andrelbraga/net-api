package service_test

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"testing"

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
	hash := srv.service.GetHash(data)
	srv.NotEmpty(hash)
}

func (srv *BookServiceStub) TestGetBookRandom() {
	user := &entities.User{Username: "test", Password: "pass"}
	hash := srv.service.GetHash(user)
	var w http.ResponseWriter
	err := srv.service.GetBookRandom(hash, w)
	if err != nil {
		srv.Error(err)
	}
}

func (srv *BookServiceStub) TestGetBookDetail() {
	user := &entities.User{Username: "test", Password: "pass"}
	hash := srv.service.GetHash(user)

	data, err := srv.service.GetBookByID(hash, "1")
	if err != nil {
		srv.Error(err)
	}

	var book *pb.Book
	err = json.Unmarshal(data, &book)
	if err != nil {
		srv.Error(err)
	}
	srv.NotEmpty(book)
}
