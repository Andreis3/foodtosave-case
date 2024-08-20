package query

import (
	"github.com/andreis3/foodtosave-case/internal/domain/usecase"
	"github.com/andreis3/foodtosave-case/internal/infra/dto"
	"github.com/andreis3/foodtosave-case/internal/infra/mapper"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type GetOneAuthorQuery struct {
	getOneAuthorService usecase.IGetOneAuthorAllBooksUseCase
}

func NewGetOneAuthorAllBooksQuery(getOneAuthorService usecase.IGetOneAuthorAllBooksUseCase) *GetOneAuthorQuery {
	return &GetOneAuthorQuery{
		getOneAuthorService: getOneAuthorService,
	}
}

func (c *GetOneAuthorQuery) Execute(id string) (dto.AuthorOutput, *util.ValidationError) {
	res, err := c.getOneAuthorService.GetOneAuthorAllBooks(id)
	if err != nil {
		return dto.AuthorOutput{}, err
	}
	return mapper.MapperAggregateToDtoOutput(res), nil
}
