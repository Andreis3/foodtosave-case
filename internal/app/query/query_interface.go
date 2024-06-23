package query

import (
	"github.com/andreis3/foodtosave-case/internal/infra/dto"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type IGetOneAuthorQuery interface {
	Execute(id string) (dto.AuthorOutput, *util.ValidationError)
}
