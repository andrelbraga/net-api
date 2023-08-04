package service

import (
	"net/http"

	"github.com/go-chi/chi"
)

func GetHash(w http.ResponseWriter, r *http.Request) {

}

func GetBookRandom(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("oi"))
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
  chi.URLParam(r, "bookId")

}
	/*
		Variables: actual_book, latest_book
		First call: latest_book = empty, actual_book: true
		Condicao para validar se os livros nao estao entre os ultimos 10 solicitados
		Validar através do hash
	*/