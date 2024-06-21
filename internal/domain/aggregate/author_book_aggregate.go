package aggregate

import "github.com/andreis3/foodtosave-case/internal/domain/entity"

type AuthorBookAggregate struct {
	Author entity.Author
	Books  []entity.Book
}

func (a AuthorBookAggregate) SetIDS(author entity.Author, books []entity.Book) AuthorBookAggregate {
	a.Author = author
	a.Books = books
	return a
}
