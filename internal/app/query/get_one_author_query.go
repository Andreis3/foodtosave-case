package query

import (
	"github.com/andreis3/foodtosave-case/internal/domain/services"
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/authorhandler/dto"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type GetAuthorQuery struct {
	getOneAuthorService services.IGetOneAuthorAllBooksService
}

func NewGetAuthorQuery(getOneAuthorService services.IGetOneAuthorAllBooksService) *GetAuthorQuery {
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
