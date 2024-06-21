package services

import (
	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type IGetAuthorService interface {
	GetOneAuthor(id string) (aggregate.AuthorBookAggregate, *util.ValidationError)
}
