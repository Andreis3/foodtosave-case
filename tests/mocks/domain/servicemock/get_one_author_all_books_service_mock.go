package servicemock

import (
	"github.com/andreis3/foodtosave-case/internal/infra/dto"
	"github.com/andreis3/foodtosave-case/internal/util"
	"github.com/stretchr/testify/mock"
)

type GetOneAuthorAllBooksServiceMock struct {
	mock.Mock
}

func (g *GetOneAuthorAllBooksServiceMock) GetOneAuthorAllBooks(id string) (dto.AuthorOutput, *util.ValidationError) {
	args := g.Called(id)
	return args.Get(0).(dto.AuthorOutput), args.Get(1).(*util.ValidationError)
}
