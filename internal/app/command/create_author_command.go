package command

import (
	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/domain/services"
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/uuid"
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/authorhandler/dto"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type CreateAuthorCommand struct {
	authorService services.ICreateAuthorWithBooksService
	uuidGenerator uuid.IUUID
}

func NewCreateAuthorCommand(authorService services.ICreateAuthorWithBooksService, uuidGenerator uuid.IUUID) *CreateAuthorCommand {
	return &CreateAuthorCommand{
		authorService: authorService,
		uuidGenerator: uuidGenerator,
	}
}

func (c *CreateAuthorCommand) Execute(data dto.AuthorInput) (dto.AuthorOutput, *util.ValidationError) {
	agg := aggregate.NewAuthorBookAggregate(c.uuidGenerator)
	agg.MapperDtoInputToAggregate(data)
	res, err := c.authorService.CreateAuthorWithBooks(*agg)
	if err != nil {
		return dto.AuthorOutput{}, err
	}
	return res, nil
}
