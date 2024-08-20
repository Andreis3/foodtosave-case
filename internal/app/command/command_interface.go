package command

import (
	"github.com/andreis3/foodtosave-case/internal/infra/dto"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type ICreateAuthorCommand interface {
	Execute(data dto.AuthorInput) (dto.AuthorOutput, *util.ValidationError)
}
