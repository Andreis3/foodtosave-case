package servicemock

import (
	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/infra/dto"
	"github.com/andreis3/foodtosave-case/internal/util"
	"github.com/stretchr/testify/mock"
)

type CreateAuthorWithBooksServiceMock struct {
	mock.Mock
}

func (m *CreateAuthorWithBooksServiceMock) CreateAuthorWithBooks(data aggregate.AuthorBookAggregate) (dto.AuthorOutput, *util.ValidationError) {
	args := m.Called(data)
	return args.Get(0).(dto.AuthorOutput), args.Get(1).(*util.ValidationError)
}
