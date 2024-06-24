package command

import (
	"github.com/andreis3/foodtosave-case/internal/domain/usecase"
	"github.com/andreis3/foodtosave-case/internal/infra/common/uuid"

	"github.com/andreis3/foodtosave-case/internal/infra/dto"
	"github.com/andreis3/foodtosave-case/internal/infra/mapper"

	"github.com/andreis3/foodtosave-case/internal/util"
)

type CreateAuthorWithBooksCommand struct {
	authorWithBooksUsecase usecase.ICreateAuthorWithBooksUsecase
	uuidGenerator          uuid.IUUID
}

func NewCreateAuthorWithBooksCommand(authorService usecase.ICreateAuthorWithBooksUsecase, uuidGenerator uuid.IUUID) *CreateAuthorWithBooksCommand {
	return &CreateAuthorWithBooksCommand{
		authorWithBooksUsecase: authorService,
		uuidGenerator:          uuidGenerator,
	}
}

func (c *CreateAuthorWithBooksCommand) Execute(data dto.AuthorInput) (dto.AuthorOutput, *util.ValidationError) {
	agg := mapper.MapperDtoInputToAggregate(data)
	res, err := c.authorWithBooksUsecase.CreateAuthorWithBooks(agg)
	if err != nil {
		return dto.AuthorOutput{}, err
	}
	return mapper.MapperAggregateToDtoOutput(res), nil
}
