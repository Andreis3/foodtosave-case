package book

import (
	"github.com/andreis3/foodtosave-case/internal/domain/entity"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type IGroupRepository interface {
	InsertBook(data entity.Book, authorId string) (*BookModel, *util.ValidationError)
	SelectAllBooksByAuthorID(authorId string) (*BookModel, *util.ValidationError)
}
