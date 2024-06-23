package services

import (
	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/authorhandler/dto"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type ICreateAuthorWithBooksService interface {
	CreateAuthorWithBooks(data aggregate.AuthorBookAggregate) (dto.AuthorOutput, *util.ValidationError)
}
type IGetOneAuthorAllBooksService interface {
	GetOneAuthorAllBooks(id string) (dto.AuthorOutput, *util.ValidationError)
}
