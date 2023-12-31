package graph

import "github.com/PauloRVF/graphql_poc_go/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AuthorDB *database.Author
	BookDB   *database.Book
}
