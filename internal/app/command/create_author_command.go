package command

import (
	"github.com/andreis3/foodtosave-case/internal/domain/services"
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/author/dto"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type CreateAuthorCommand struct {
	authorService services.ICreateAuthorService
	output        dto.AuthorOutput
}

func NewCreateGroupCommand(authorService services.ICreateAuthorService) *CreateAuthorCommand {
	return &CreateAuthorCommand{
		authorService: authorService,
	}
}
func (c *CreateAuthorCommand) Execute(data dto.AuthorInput) (dto.AuthorOutput, *util.ValidationError) {
	aggregate := data.MapperToAggregateAuthor()
	res, err := c.authorService.CreateAuthor(aggregate)
	if err != nil {
		return dto.AuthorOutput{}, err
	}
	return c.output.MapperToAggregateAuthor(res), nil
}
