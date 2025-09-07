package store

import (
	"aprende-golang/internal/model"
	"database/sql"
)

type Store interface {
	GetAll() ([]*model.Book, error)
	GetByID(id int) (*model.Book, error)
	Create(book *model.Book) (*model.Book, error)
	Update(id int, book *model.Book) (*model.Book, error)
	Delete(id int) error
}

type store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return &store{db: db}
}

func (s *store) GetAll() ([]*model.Book, error) {
	q := `SELECT id, title, author FROM books`
	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []*model.Book

	for rows.Next() {
		var b = &model.Book{}
		if err := rows.Scan(&b.Id, &b.Title, &b.Author); err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	return books, nil
}

func (s *store) GetByID(Id int) (*model.Book, error) {
	q := `SELECT id, title, author FROM books WHERE id = ?`
	var b = &model.Book{}

	err := s.db.QueryRow(q, Id).Scan(&b.Id, &b.Title, &b.Author)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (s *store) Create(book *model.Book) (*model.Book, error) {
	q := `INSERT INTO books (title, author) VALUES (?, ?)`

	resp, err := s.db.Exec(q, book.Title, book.Author)
	if err != nil {
		return nil, err
	}

	id, err := resp.LastInsertId()
	if err != nil {
		return nil, err
	}

	book.Id = int(id)

	return book, nil

}

func (s *store) Update(id int, book *model.Book) (*model.Book, error) {
	q := `UPDATE books SET title = ?, author = ? WHERE id = ?`

	_, err := s.db.Exec(q, book.Title, book.Author, id)
	if err != nil {
		return nil, err
	}

	book.Id = id

	return book, nil
}

func (s *store) Delete(id int) error {
	q := `DELETE FROM books WHERE id = ?`

	_, err := s.db.Exec(q, id)
	if err != nil {
		return err
	}
	return nil
}
