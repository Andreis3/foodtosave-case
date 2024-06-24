package query

import (
	"github.com/andreis3/foodtosave-case/internal/domain/usecase"
	"github.com/andreis3/foodtosave-case/internal/infra/dto"
	"github.com/andreis3/foodtosave-case/internal/infra/mapper"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type GetOneAuthorAllBooksQuery struct {
	getOneAuthorService usecase.IGetOneAuthorAllBooksUsecase
}

func NewGetOneAuthorAllBooksQuery(getOneAuthorService usecase.IGetOneAuthorAllBooksUsecase) *GetOneAuthorAllBooksQuery {
	return &GetOneAuthorAllBooksQuery{
		getOneAuthorService: getOneAuthorService,
	}
}

func (c *GetOneAuthorAllBooksQuery) Execute(id string) (dto.AuthorOutput, *util.ValidationError) {
	res, err := c.getOneAuthorService.GetOneAuthorAllBooks(id)
	if err != nil {
		return dto.AuthorOutput{}, err
	}
	return mapper.MapperAggregateToDtoOutput(res), nil
}
