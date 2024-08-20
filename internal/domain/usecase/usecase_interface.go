package usecase

import (
	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type ICreateAuthorWithBooksUseCase interface {
	CreateAuthorWithBooks(data aggregate.AuthorBookAggregate) (aggregate.AuthorBookAggregate, *util.ValidationError)
}
type IGetOneAuthorAllBooksUseCase interface {
	GetOneAuthorAllBooks(id string) (aggregate.AuthorBookAggregate, *util.ValidationError)
}
