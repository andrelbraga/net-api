package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"net-api.com/internal/domain/entities"
)

type BookServiceInterface interface {
	GetBookByID(hash string, bookIdParam string) ([]byte, error)
	GetBookRandom(hash string, w http.ResponseWriter) error
	GetHash(user *entities.User) string
}

// BooksController
type BooksController struct {
	BookServiceInterface
}

// NewBooksController
func NewBooksController(srv BookServiceInterface) *BooksController {
	return &BooksController{
		srv,
	}
}

// Post user entities.User godoc
// @Summary Post user Url
// @Description Post user Url
// @Accept  json
// @Produce  json
// @Param login body entities.User true "login"
// @Success 200 {object} controller.HTTPResponse
// @Failure 500 {object} controller.HTTPResponse
// @Router /user/hash [post]
func (srv *BooksController) GetHashCtrl(w http.ResponseWriter, r *http.Request) {
	user := &entities.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h := srv.GetHash(user)
	body := &HTTPBody{
		Message: h,
	}
	var response = HTTPResponse{
		StatusCode: 200,
		Body:       body,
		Error:      "",
	}
	data, _ := json.Marshal(response)
	w.Write(data)
}

// Get random entities.Book godoc
// @Summary Get books Url
// @Description Get random books Url
// @Accept  json
// @Produce  json
// @Param api-key header string true "Api Key"
// @Success 200 {object} controller.HTTPResponse
// @Failure 500 {object} controller.HTTPResponse
// @Router /book/random [get]
func (srv *BooksController) GetBookRandomCtrl(w http.ResponseWriter, r *http.Request) {
	hash := r.Header.Get("Api-Key")
	if hash == "" {
		http.Error(w, "Hash not found at header", http.StatusBadRequest)
		return
	}

	err := srv.GetBookRandom(hash, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// Book entities.Book godoc
// @Summary Get book by id
// @Description Get book by id
// @Accept  json
// @Produce  json
// @Param api-key header string true "Api Key"
// @Param id path int true "Book ID"
// @Success 200 {object} controller.HTTPResponse
// @Failure 500 {object} controller.HTTPResponse
// @Router /book/{id} [get]
func (srv *BooksController) GetBookByIDCtrl(w http.ResponseWriter, r *http.Request) {
	hash := r.Header.Get("Api-Key")
	if hash == "" {
		http.Error(w, "Hash not found at header", http.StatusBadRequest)
		return
	}

	bookIdParam := chi.URLParam(r, "id")
	if bookIdParam == "" {
		http.Error(w, "param {id} not found", http.StatusBadRequest)
		return
	}

	data, err := srv.GetBookByID(hash, bookIdParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(data)
}
