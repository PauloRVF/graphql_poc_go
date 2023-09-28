package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Book struct {
	db              *sql.DB
	ID              string
	Title           string
	AuthorID        string
	PublicationYear int
	Rating          sql.NullInt64
}

func NewBook(db *sql.DB) *Book {
	return &Book{db: db}
}

func (a *Book) Create(title string, authorID string, publicationYear int) (Book, error) {
	id := uuid.New().String()
	_, err := a.db.Exec("INSERT INTO books(id, title, author_id, publication_year) VALUES ($1, $2, $3, $4)", id, title, authorID, publicationYear)
	if err != nil {
		return Book{}, err
	}
	return Book{ID: id, Title: title, AuthorID: authorID, PublicationYear: publicationYear}, nil
}

func (a *Book) FindAll() ([]Book, error) {
	rows, err := a.db.Query("SELECT id, title, author_id, publication_year, rating FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []Book{}
	for rows.Next() {
		var id, title, authorID string
		var publicationYear int
		var rating sql.NullInt64
		err := rows.Scan(&id, &title, &authorID, &publicationYear, &rating)
		if err != nil {
			return nil, err
		}
		books = append(books, Book{ID: id, Title: title, AuthorID: authorID, PublicationYear: publicationYear, Rating: rating})
	}

	return books, nil
}

func (a *Book) FindByAuthor(authorId string) ([]Book, error) {
	rows, err := a.db.Query("SELECT id, title, author_id, publication_year, rating FROM books WHERE author_id = $1", authorId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []Book{}
	for rows.Next() {
		var id, title, authorID string
		var publicationYear int
		var rating sql.NullInt64
		err := rows.Scan(&id, &title, &authorID, &publicationYear, &rating)
		if err != nil {
			return nil, err
		}
		books = append(books, Book{ID: id, Title: title, AuthorID: authorID, PublicationYear: publicationYear, Rating: rating})
	}

	return books, nil
}
