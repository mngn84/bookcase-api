package store

import (
	bookmodel "github.com/mngn84/bookcase-api/internal/app/models"
)

type BookRepository interface {
	AddBook(*bookmodel.FullBookRequest, *bookmodel.Book) error
	GetAllBooks() ([]*bookmodel.Book, error)
	FindAllBooksByAuthor(string) ([]*bookmodel.Book, error)
	FindBookByID(int) (*bookmodel.Book, error)
	MoveBook(int, *bookmodel.Location) error
	MarkAsRead(int) error
	RemoveBook(int) error
}