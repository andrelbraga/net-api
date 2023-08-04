package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net-api.com/internal/controller"
)

func main() {
	r := chi.NewRouter()
	
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	controller.RegisterRoutes(r)

	http.ListenAndServe(":3000", r)
}
