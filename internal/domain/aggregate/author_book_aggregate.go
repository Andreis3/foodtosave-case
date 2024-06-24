package aggregate

import (
	"github.com/andreis3/foodtosave-case/internal/domain/entity"
	"github.com/andreis3/foodtosave-case/internal/domain/errors"
)

type AuthorBookAggregate struct {
	Author entity.Author
	Books  []entity.Book
}

func NewAuthorBookAggregate() *AuthorBookAggregate {
	return &AuthorBookAggregate{
		Author: entity.Author{},
		Books:  []entity.Book{},
	}
}
func (a *AuthorBookAggregate) AddAuthorAndBooks(author entity.Author, books []entity.Book) {
	a.Author = author
	a.Books = books
}

func (a *AuthorBookAggregate) Validate() *errors.NotificationErrors {
	authorValidate := a.Author.Validate()
	if len(a.Books) == 0 {
		authorValidate.AddErrors("books: minimum 1 book is required")
	}
	for index := range a.Books {
		authorValidate.MergeErrors(index, "books", a.Books[index].Validate())
	}
	return authorValidate
}
