package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Author struct {
	db   *sql.DB
	ID   string
	Name string
}

func NewAuthor(db *sql.DB) *Author {
	return &Author{db: db}
}

func (a *Author) Create(name string) (Author, error) {
	id := uuid.New().String()
	_, err := a.db.Exec("INSERT INTO authors (id, name) VALUES ($1, $2)", id, name)
	if err != nil {
		return Author{}, err
	}
	return Author{db: a.db, ID: id, Name: name}, nil
}

func (a *Author) FindbyBookID(BookID string) (Author, error) {
	var id, name string
	err := a.db.QueryRow("SELECT a.id, a.name FROM authors a JOIN books b ON a.id = b.author_id WHERE b.id = $1", BookID).Scan(&id, &name)
	if err != nil {
		return Author{}, err
	}

	return Author{ID: id, Name: name}, nil
}

func (a *Author) FindAll() ([]Author, error) {
	rows, err := a.db.Query("SELECT id, name FROM authors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	authors := []Author{}
	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		authors = append(authors, Author{ID: id, Name: name})
	}
	return authors, nil
}
