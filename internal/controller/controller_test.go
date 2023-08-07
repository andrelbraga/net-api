package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"net-api.com/internal/controller"
	mock_controller "net-api.com/internal/controller/mock"
	pb "net-api.com/internal/infra/grpc/proto"

	"net-api.com/internal/domain/entities"
)

type BookControllerStub struct {
	suite.Suite
	mockSrv *mock_controller.MockBookServiceInterface
	ctrl    *controller.BooksController
}

func TestBookController(t *testing.T) {
	suite.Run(t, new(BookControllerStub))
}

func (s *BookControllerStub) SetupSuite() {
	mockCtrl := gomock.NewController(s.T())
	defer mockCtrl.Finish()

	s.mockSrv = mock_controller.NewMockBookServiceInterface(mockCtrl)
	s.ctrl = controller.NewBooksController(s.mockSrv)
}

func (stub *BookControllerStub) TestGetHashCtrl() {
	user := &entities.User{Username: "any_username", Password: "any_password"}
	body, _ := json.Marshal(user)
	hash := "any_hash"

	// Mock
	stub.mockSrv.EXPECT().GetHash(user).Return(hash).Times(1)

	req, _ := http.NewRequest("POST", "http://localhost:9191/api/v1/user/hash", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	stub.ctrl.GetHashCtrl(rr, req)

	// Assertions
	stub.Equal(http.StatusOK, rr.Code)

	var response controller.HTTPResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		stub.Errorf(err, "error decoding response body: %v")
	}
	stub.Equal(response.Body.Message, hash)
}

func (stub *BookControllerStub) TestGetBookByIDCtrl() {
	hash := "any_hash"
	paramId := "1"

	req, _ := http.NewRequest("GET", "http://localhost:9191/api/v1/book/1", nil)
	req.Header.Set("Api-Key", hash)
	rr := httptest.NewRecorder()

	book := createBookDetail(paramId)
	data, _ := json.Marshal(book.Book)

	// Mock
	stub.mockSrv.EXPECT().GetBookByID(hash, paramId).Return(data, nil).Times(1)

	stub.ctrl.GetBookByIDCtrl(rr, req)

	// Assertions
	stub.Equal(http.StatusBadRequest, rr.Code)
}

func createBookDetail(id string) *pb.GetBookDetailsResponse {
	return &pb.GetBookDetailsResponse{
		BookId: id,
		Book: &pb.Book{
			Id:            id,
			Title:         "any_title",
			Authors:       []string{},
			ImageLinks:    nil,
			PrintType:     "BOOK",
			Language:      "pt-br",
			PublishedDate: nil,
			PageCount:     nil,
			Description:   nil,
		},
	}
}
