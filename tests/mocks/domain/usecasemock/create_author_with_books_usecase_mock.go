package usecasemock

import (
	"github.com/stretchr/testify/mock"

	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type CreateAuthorWithBooksUseCaseMock struct {
	mock.Mock
}

func (m *CreateAuthorWithBooksUseCaseMock) CreateAuthorWithBooks(data aggregate.AuthorBookAggregate) (aggregate.AuthorBookAggregate, *util.ValidationError) {
	args := m.Called(data)
	return args.Get(0).(aggregate.AuthorBookAggregate), args.Get(1).(*util.ValidationError)
}
