package command

import (
	"github.com/andreis3/foodtosave-case/internal/domain/usecase"
	"github.com/andreis3/foodtosave-case/internal/infra/common/uuid"

	"github.com/andreis3/foodtosave-case/internal/infra/dto"
	"github.com/andreis3/foodtosave-case/internal/infra/mapper"

	"github.com/andreis3/foodtosave-case/internal/util"
)

type CreateAuthorCommand struct {
	authorWithBooksUseCase usecase.ICreateAuthorWithBooksUseCase
	uuidGenerator          uuid.IUUID
}

func NewCreateAuthorWithBooksCommand(authorService usecase.ICreateAuthorWithBooksUseCase, uuidGenerator uuid.IUUID) *CreateAuthorCommand {
	return &CreateAuthorCommand{
		authorWithBooksUseCase: authorService,
		uuidGenerator:          uuidGenerator,
	}
}

func (c *CreateAuthorCommand) Execute(data dto.AuthorInput) (dto.AuthorOutput, *util.ValidationError) {
	agg := mapper.MapperDtoInputToAggregate(data)
	res, err := c.authorWithBooksUseCase.CreateAuthorWithBooks(agg)
	if err != nil {
		return dto.AuthorOutput{}, err
	}
	return mapper.MapperAggregateToDtoOutput(res), nil
}
