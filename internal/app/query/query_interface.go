package query

import (
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/authorhandler/dto"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type IGetOneAuthorQuery interface {
	Execute(id string) (dto.AuthorOutput, *util.ValidationError)
}
