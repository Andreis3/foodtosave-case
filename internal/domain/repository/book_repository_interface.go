package repository

import (
	"github.com/andreis3/foodtosave-case/internal/domain/entity"
	"github.com/andreis3/foodtosave-case/internal/infra/repository/postgres/book"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type IBookRepository interface {
	InsertBook(data entity.Book, authorId string) (string, *util.ValidationError)
	SelectAllBooksByAuthorID(authorId string) ([]book.BookModel, *util.ValidationError)
}
