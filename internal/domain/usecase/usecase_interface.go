package usecase

import (
	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type ICreateAuthorWithBooksUsecase interface {
	CreateAuthorWithBooks(data aggregate.AuthorBookAggregate) (aggregate.AuthorBookAggregate, *util.ValidationError)
}
type IGetOneAuthorAllBooksUsecase interface {
	GetOneAuthorAllBooks(id string) (aggregate.AuthorBookAggregate, *util.ValidationError)
}
