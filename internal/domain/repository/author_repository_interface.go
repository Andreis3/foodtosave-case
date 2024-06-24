package repository

import (
	"github.com/andreis3/foodtosave-case/internal/domain/entity"
	"github.com/andreis3/foodtosave-case/internal/infra/repository/postgres/author"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type IAuthorRepository interface {
	InsertAuthor(data entity.Author) (string, *util.ValidationError)
	SelectOneAuthorByID(authorId string) (*author.AuthorModel, *util.ValidationError)
}
