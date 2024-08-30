package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mngn84/bookcase-api/internal/app/handlers"
	bookmodel "github.com/mngn84/bookcase-api/internal/app/models"
	"github.com/mngn84/bookcase-api/internal/app/store"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	r := s.router
	bh := &handlers.BookHandler{Store: s.store.Book()}

	r.HandleFunc("/books", bh.HandleAddBook(&bookmodel.FullBookRequest{})).Methods("POST")
	r.HandleFunc("/books", bh.HandleGetAllBooks()).Methods("GET")
	r.HandleFunc("/books/{author}", bh.HandleFindAllBooksByAuthor()).Methods("GET")
	r.HandleFunc("/books/{id}", bh.HandleFindBookByID()).Methods("GET")
	r.HandleFunc("/books/{id}/move", bh.HandleMoveBook(&bookmodel.MoveBookRequest{})).Methods("PUT")
	r.HandleFunc("/books/{id}/read", bh.HandleMarkAsRead()).Methods("PUT")
	r.HandleFunc("/books/{id}", bh.HandleRemoveBook()).Methods("DELETE")
}
