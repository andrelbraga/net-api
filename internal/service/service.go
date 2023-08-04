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
	"github.com/go-chi/render"
	"google.golang.org/protobuf/types/known/emptypb"

	"net-api.com/internal/domain/entities"
	"net-api.com/internal/infra/grpc"

	pb "net-api.com/internal/infra/grpc/proto"
)

func GetHash(w http.ResponseWriter, r *http.Request) {
	data := &entities.UserRequest{}
	if err := render.Bind(r, data); err != nil {
		w.Write([]byte(err.Error()))
	}

	s := fmt.Sprintf("%s:%s", data.Login, data.Password)
	h := GetMD5Hash(s)
	w.Write([]byte(h))

	fmt.Println(s, h)
}

func GetBookRandom(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.NewBookClient()
	if err != nil {
		log.Print(err.Error())
	}

	client := pb.NewPrivateBookServiceClient(conn)
	empty := &emptypb.Empty{}
	stream, err := client.GetRandomBook(context.Background(), empty)
	defer conn.Close()

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	for {
		book, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListBooks(_) = _, %v", client, err)
		}
		log.Println(book)
		data, _ := json.Marshal(book)
		w.Write(data)
	}
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	bookIdParam := chi.URLParam(r, "bookId")
	conn, err := grpc.NewBookClient()
	if err != nil {
		log.Print(err.Error())
	}

	b := pb.NewPrivateBookServiceClient(conn)
	book, err := b.GetBookDetail(context.Background(), &pb.GetBookDetailsRequest{BookId: bookIdParam})
	defer conn.Close()

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	data, _ := json.Marshal(book.Book)
	w.Write(data)
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GetMock(w http.ResponseWriter, r *http.Request) {
	var book []map[string]interface{}

	for i := 0; i < 10; i++ {
		item := make(map[string]interface{})
		item["id"] = i
		item["name"] = fmt.Sprintf("Item %d", i)
		book = append(book, item)
	}
	data, _ := json.Marshal(book)

	w.Write([]byte(data))
}

/*
	Variables: actual_book, latest_book
	First call: latest_book = empty, actual_book: true
	Condicao para validar se os livros nao estao entre os ultimos 10 solicitados
	Validar através do hash
*/
