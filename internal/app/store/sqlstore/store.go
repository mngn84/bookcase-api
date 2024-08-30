package sqlstore

import (
	"database/sql"
	"github.com/mngn84/bookcase-api/internal/app/store"

	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	db             *sql.DB
	bookRepository *BookRepository
}

// New ...
func New(db *sql.DB) *Store {
	st := &Store{
		db: db,
	}
	st.bookRepository = &BookRepository{
		store: st,
	}
	return st
}

// Book ...
func (st *Store) Book() store.BookRepository {
	return st.bookRepository
}
