package sqlstore

import (
	bookmodel "github.com/mngn84/bookcase-api/internal/app/models"
)

// BookRepository ...
type BookRepository struct {
	store *Store
}

// AddBook ...
func (r *BookRepository) AddBook(fr *bookmodel.FullBookRequest, b *bookmodel.Book) error {
	tr, err := r.store.db.Begin()
	if err != nil {
		return err
	}
	defer tr.Rollback()

	var locId int

	err = tr.QueryRow(
		"INSERT INTO locations (case_id, shelf_name) VALUES ($1, $2) RETURNING id",
		fr.Location.CaseId, fr.Location.ShelfName,
	).Scan(&locId)
	if err != nil {
		return err
	}
	var bId int
	err = tr.QueryRow(
		"INSERT INTO books (title, author, genre, is_read, location_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		b.Title, b.Author, b.Genre, b.IsRead, locId,
	).Scan(
		&bId,
	); if err != nil {
		return err
	}
	_, err = tr.Exec("UPDATE locations SET book_id = $1 WHERE id = $2", bId, locId)
	
	if err != nil {
		return err
	}

	return tr.Commit()
}

// GetAllBooks ...
func (r *BookRepository) GetAllBooks() ([]*bookmodel.Book, error) {
	rows, err := r.store.db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*bookmodel.Book
	for rows.Next() {
		b := &bookmodel.Book{}
		if err := rows.Scan(
			&b.ID,
			&b.Title,
			&b.Author,
			&b.Genre,
			&b.IsRead,
			&b.LocationId,
		); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}
 
//FindAllBooksByAuthor ...
func (r *BookRepository) FindAllBooksByAuthor(author string) ([]*bookmodel.Book, error) {
	rows, err := r.store.db.Query("SELECT * FROM books WHERE author = $1", author)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*bookmodel.Book
	for rows.Next() {
		b := &bookmodel.Book{}
		if err := rows.Scan(
			&b.ID,
			&b.Title,
			&b.Author,
			&b.Genre,
			&b.IsRead,
			&b.LocationId,
		); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

// FindBookByID ...
func (r *BookRepository) FindBookByID(id int) (*bookmodel.Book, error) {
	book := &bookmodel.Book{}
	if err := r.store.db.QueryRow(
		"SELECT * FROM books WHERE id = $1",
		id,
	).Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.Genre,
		&book.IsRead,
		&book.LocationId,
	); err != nil {
		return nil, err
	}
	return book, nil
}

// MoveBook ...
func (r *BookRepository) MoveBook(id int, loc *bookmodel.Location) error {
	b := &bookmodel.Book{}
	if err := r.store.db.QueryRow(
		"UPDATE locations l SET case_id = $1, shelf_name = $2 FROM books b WHERE b.location_id = l.id AND b.id = $3 RETURNING b.id",
		loc.CaseId, loc.ShelfName, id,
	).Scan(
		&b.ID,
	); err != nil {
		return err
	}
	return nil
}

// MarkAsRead ...
func (r *BookRepository) MarkAsRead(id int) error {
	b := &bookmodel.Book{}
	if err := r.store.db.QueryRow(
		"UPDATE books SET is_read = true WHERE id = $1 RETURNING id",
		id,
	).Scan(
		&b.ID,
	); err != nil {
		return err
	}
	return nil
}

// RemoveBook ...
func (r *BookRepository) RemoveBook(id int) error {
	b := &bookmodel.Book{}
	if err := r.store.db.QueryRow(
		"DELETE FROM books WHERE id = $1 RETURNING id",
		id,
	).Scan(
		&b.ID,
	); err != nil {
		return err
	}
	return nil
}
