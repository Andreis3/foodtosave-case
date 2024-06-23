package command

import (
	dto2 "github.com/andreis3/foodtosave-case/internal/infra/dto"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type ICreateAuthorCommand interface {
	Execute(data dto2.AuthorInput) (dto2.AuthorOutput, *util.ValidationError)
}
