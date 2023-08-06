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

	"github.com/go-chi/chi"

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
func (bsrv *BookService) GetHash(w http.ResponseWriter, r *http.Request) {
	user := &entities.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s := fmt.Sprintf("%s:%s", user.Username, user.Password)
	h := getMD5Hash(s)
	w.Write([]byte(h))
}

// GetBookRandom
func (bsrv *BookService) GetBookRandom(w http.ResponseWriter, r *http.Request) {
	hash := r.Header.Get("Api-Key")

	stream, err := bsrv.client.GetRandomBook(context.Background(), &pb.GetBookRandomRequest{ApiKey: hash})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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
}

// GetBookByID
func (bsrv *BookService) GetBookByID(w http.ResponseWriter, r *http.Request) {
	bookIdParam := chi.URLParam(r, "bookId")
	hash := r.Header.Get("Api-Key")

	book, err := bsrv.client.GetBookDetail(context.Background(), &pb.GetBookDetailsRequest{ApiKey: hash, BookId: bookIdParam})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, _ := json.Marshal(book.Book)
	w.Write(data)
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
