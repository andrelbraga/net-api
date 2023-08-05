package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"net-api.com/internal/controller"
	"net-api.com/internal/infra/grpc"
	pb "net-api.com/internal/infra/grpc/proto"
	"net-api.com/internal/service"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	conn, err := grpc.NewBookClient()
	if err != nil {
		log.Print(err.Error())
	}

	client := pb.NewPrivateBookServiceClient(conn)

	srv := service.NewBookService(client)
	ctrl := controller.NewBooksController(srv)
	ctrl.RegisterRoutes(r)

	http.ListenAndServe(":9191", r)
}
