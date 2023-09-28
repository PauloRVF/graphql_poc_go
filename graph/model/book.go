package model

type Book struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	PublicationYear int    `json:"publication_year"`
	Rating          *int   `json:"rating,omitempty"`
}
