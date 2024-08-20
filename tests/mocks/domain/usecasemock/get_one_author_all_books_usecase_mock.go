package usecasemock

import (
	"github.com/stretchr/testify/mock"

	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type GetOneAuthorAllBooksUseCaseMock struct {
	mock.Mock
}

func (g *GetOneAuthorAllBooksUseCaseMock) GetOneAuthorAllBooks(id string) (aggregate.AuthorBookAggregate, *util.ValidationError) {
	args := g.Called(id)
	return args.Get(0).(aggregate.AuthorBookAggregate), args.Get(1).(*util.ValidationError)
}
