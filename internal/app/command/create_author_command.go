package command

import (
	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/domain/usecase"
	"github.com/andreis3/foodtosave-case/internal/infra/common/uuid"
	dto2 "github.com/andreis3/foodtosave-case/internal/infra/dto"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type CreateAuthorCommand struct {
	authorService usecase.ICreateAuthorWithBooksService
	uuidGenerator uuid.IUUID
}

func NewCreateAuthorCommand(authorService usecase.ICreateAuthorWithBooksService, uuidGenerator uuid.IUUID) *CreateAuthorCommand {
	return &CreateAuthorCommand{
		authorService: authorService,
		uuidGenerator: uuidGenerator,
	}
}

func (c *CreateAuthorCommand) Execute(data dto2.AuthorInput) (dto2.AuthorOutput, *util.ValidationError) {
	agg := aggregate.NewAuthorBookAggregate(c.uuidGenerator)
	agg.MapperDtoInputToAggregate(data)
	res, err := c.authorService.CreateAuthorWithBooks(*agg)
	if err != nil {
		return dto2.AuthorOutput{}, err
	}
	return res, nil
}
