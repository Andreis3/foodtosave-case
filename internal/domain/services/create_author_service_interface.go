package services

import (
	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type ICreateAuthorService interface {
	CreateAuthor(data aggregate.AuthorBookAggregate) (aggregate.AuthorBookAggregate, *util.ValidationError)
}
