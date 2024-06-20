package author

import (
	"github.com/andreis3/foodtosave-case/internal/domain/entity"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type IAuthorRepository interface {
	InsertAuthor(data entity.Author) (*AuthorModel, *util.ValidationError)
	SelectOneAuthorByID(authorId string) (*AuthorModel, *util.ValidationError)
}
