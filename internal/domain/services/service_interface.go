package services

import (
	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/author/dto"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type ICreateAuthorService interface {
	CreateAuthor(data aggregate.AuthorBookAggregate) (dto.AuthorOutput, *util.ValidationError)
}
type IGetAuthorService interface {
	GetOneAuthor(id string) (dto.AuthorOutput, *util.ValidationError)
}
