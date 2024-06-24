package usecasemock

import (
	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/util"
	"github.com/stretchr/testify/mock"
)

type CreateAuthorWithBooksUsecaseMock struct {
	mock.Mock
}

func (m *CreateAuthorWithBooksUsecaseMock) CreateAuthorWithBooks(data aggregate.AuthorBookAggregate) (aggregate.AuthorBookAggregate, *util.ValidationError) {
	args := m.Called(data)
	return args.Get(0).(aggregate.AuthorBookAggregate), args.Get(1).(*util.ValidationError)
}
