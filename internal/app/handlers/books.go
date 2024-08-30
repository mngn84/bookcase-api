package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	bookmodel "github.com/mngn84/bookcase-api/internal/app/models"
	"github.com/mngn84/bookcase-api/internal/app/store"
)

type BookHandler struct {
	Store store.BookRepository
}

// HandleAddBook...
func (h *BookHandler) HandleAddBook(fr *bookmodel.FullBookRequest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(fr); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		b := &bookmodel.Book{
			Title:  fr.Title,
			Author: fr.Author,
			Genre:  fr.Genre,
			IsRead: fr.IsRead,
		}

		if err := h.Store.AddBook(fr, b); err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}

		h.respond(w, r, http.StatusCreated, b)

	}
}

// HandleGetAllBooks...
func (h *BookHandler) HandleGetAllBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := h.Store.GetAllBooks()
		if err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}
		h.respond(w, r, http.StatusOK, books)
	}
}

// HandleFindBookByAuthor...
func (h *BookHandler) HandleFindAllBooksByAuthor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		author := vars["author"]
		books, err := h.Store.FindAllBooksByAuthor(author)
		if err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}
		h.respond(w, r, http.StatusOK, books)
	}
}

// HandleFindBookByID...
func (h *BookHandler) HandleFindBookByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			return
		}

		b, err := h.Store.FindBookByID(id)

		if err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}

		h.respond(w, r, http.StatusOK, b)
	}
}

// HandleMoveBook...
func (h *BookHandler) HandleMoveBook(mr *bookmodel.MoveBookRequest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			return
		}
		if err := json.NewDecoder(r.Body).Decode(mr); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.Store.MoveBook(id, &mr.Location); err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}
		h.respond(w, r, http.StatusOK, nil)
	}
}

// HandleMarkAsRead...
func (h *BookHandler) HandleMarkAsRead() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			return
		}
		if err := h.Store.MarkAsRead(id); err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}
		h.respond(w, r, http.StatusOK, nil)
	}
}

// HandleRemoveBook...
func (h *BookHandler) HandleRemoveBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			return
		}
		if err := h.Store.RemoveBook(id); err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}
		h.respond(w, r, http.StatusOK, nil)
	}
}

func (h *BookHandler) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	h.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (h *BookHandler) respond(w http.ResponseWriter, _ *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
