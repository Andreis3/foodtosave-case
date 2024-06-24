package usecase

import (
	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/infra/dto"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type ICreateAuthorWithBooksUsecase interface {
	CreateAuthorWithBooks(data aggregate.AuthorBookAggregate) (dto.AuthorOutput, *util.ValidationError)
}
type IGetOneAuthorAllBooksUsecase interface {
	GetOneAuthorAllBooks(id string) (dto.AuthorOutput, *util.ValidationError)
}
