package query

import (
	"github.com/andreis3/foodtosave-case/internal/domain/usecase"
	"github.com/andreis3/foodtosave-case/internal/infra/dto"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type GetAuthorQuery struct {
	getOneAuthorService usecase.IGetOneAuthorAllBooksService
}

func NewGetAuthorQuery(getOneAuthorService usecase.IGetOneAuthorAllBooksService) *GetAuthorQuery {
	return &GetAuthorQuery{
		getOneAuthorService: getOneAuthorService,
	}
}

func (c *GetAuthorQuery) Execute(id string) (dto.AuthorOutput, *util.ValidationError) {
	res, err := c.getOneAuthorService.GetOneAuthorAllBooks(id)
	if err != nil {
		return dto.AuthorOutput{}, err
	}
	return res, nil
}
