package command

import (
	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/domain/usecase"
	"github.com/andreis3/foodtosave-case/internal/domain/uuid"
	"github.com/andreis3/foodtosave-case/internal/infra/dto"

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
	agg := aggregate.NewAuthorBookAggregate(c.uuidGenerator)
	agg.MapperDtoInputToAggregate(data)
	res, err := c.authorWithBooksUsecase.CreateAuthorWithBooks(*agg)
	if err != nil {
		return dto.AuthorOutput{}, err
	}
	return res, nil
}
