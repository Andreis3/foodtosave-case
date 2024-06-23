package query

import (
	"github.com/andreis3/foodtosave-case/internal/domain/services"
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/author/dto"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type GetAuthorQuery struct {
	getOneAuthorService services.IGetAuthorService
}

func NewGetAuthorQuery(getOneAuthorService services.IGetAuthorService) *GetAuthorQuery {
	return &GetAuthorQuery{
		getOneAuthorService: getOneAuthorService,
	}
}

func (c *GetAuthorQuery) Execute(id string) (dto.AuthorOutput, *util.ValidationError) {
	res, err := c.getOneAuthorService.GetOneAuthor(id)
	if err != nil {
		return dto.AuthorOutput{}, err
	}
	return res, nil
}
