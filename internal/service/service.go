package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"net-api.com/internal/domain/entities"

	pb "net-api.com/internal/infra/grpc/proto"
)

// BookService
type BookService struct {
	client pb.PrivateBookServiceClient
}

// NewBookService
func NewBookService(client pb.PrivateBookServiceClient) *BookService {
	return &BookService{
		client: client,
	}
}

// GetHash
func (bsrv *BookService) GetHash(user *entities.User) string {
	s := fmt.Sprintf("%s:%s", user.Username, user.Password)
	h := getMD5Hash(s)
	return h
}

// GetBookRandom
func (bsrv *BookService) GetBookRandom(hash string, w http.ResponseWriter) error {

	stream, err := bsrv.client.GetRandomBook(context.Background(), &pb.GetBookRandomRequest{ApiKey: hash})
	if err != nil {
		return err
	}

	for {
		book, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListBooks(_) = _, %v", bsrv.client, err)
		}
		log.Println(book)
		data, _ := json.Marshal(book)
		w.Write(data)
	}
	return nil
}

// GetBookByID
func (bsrv *BookService) GetBookByID(hash, bookIdParam string) ([]byte, error) {
	book, err := bsrv.client.GetBookDetail(context.Background(), &pb.GetBookDetailsRequest{ApiKey: hash, BookId: bookIdParam})
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(book.Book)
	return data, nil
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
