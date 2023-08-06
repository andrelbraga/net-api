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
	"net-api.com/internal/router"
	"net-api.com/internal/service"
)

// @title Echo Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9191
// @BasePath /api/v1
// @schemes http
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
	rtr := router.NewBooksRouter(ctrl)
	rtr.RegisterRoutes(r)

	http.ListenAndServe(":9191", r)
}
