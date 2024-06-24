//go:build unit
// +build unit

package command_test

import (
	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/domain/entity"
	"github.com/andreis3/foodtosave-case/internal/infra/dto"
	"github.com/andreis3/foodtosave-case/internal/infra/mapper"
	"github.com/andreis3/foodtosave-case/internal/util"
	"github.com/andreis3/foodtosave-case/tests/mocks/domain/usecasemock"
	"github.com/andreis3/foodtosave-case/tests/mocks/infra/common/uuidmock"
	"github.com/stretchr/testify/mock"
)

func ContextCreateSuccess() (*usecasemock.CreateAuthorWithBooksUsecaseMock, *uuidmock.UUIDMock, aggregate.AuthorBookAggregate) {
	authorWithBooksUsecaseMock := new(usecasemock.CreateAuthorWithBooksUsecaseMock)
	uuidMock := new(uuidmock.UUIDMock)
	output := aggregate.AuthorBookAggregate{
		Author: entity.Author{
			ID:          "1",
			Name:        "Author 1",
			Nationality: "Brazilian",
		},
		Books: []entity.Book{
			{
				ID:     "1",
				Title:  "Book 1",
				Gender: "Terror",
			},
			{
				ID:     "2",
				Title:  "Book 2",
				Gender: "Comedy",
			},
		},
	}

	authorWithBooksUsecaseMock.On("CreateAuthorWithBooks", mock.Anything).Return(output, (*util.ValidationError)(nil))
	uuidMock.On("Generate").Return("111")

	return authorWithBooksUsecaseMock, uuidMock, output
}

func ContextCreateError() (*usecasemock.CreateAuthorWithBooksUsecaseMock, *uuidmock.UUIDMock, *util.ValidationError) {
	authorWithBooksUsecaseMock := new(usecasemock.CreateAuthorWithBooksUsecaseMock)
	uuidMock := new(uuidmock.UUIDMock)
	output := aggregate.AuthorBookAggregate{}

	uuidMock.On("Generate").Return("111")

	callExpected := mapper.MapperDtoInputToAggregate(dto.AuthorInput{
		Name:        "Author 1",
		Nationality: "Brazilian",
		Books: []dto.BookInput{
			{
				Title:  "Book 1",
				Gender: "Terror",
			},
			{
				Title:  "Book 2",
				Gender: "Comedy",
			},
		},
	})

	err := &util.ValidationError{
		Code:        "PIDB-235",
		Status:      500,
		ClientError: []string{"Internal Server Error"},
		LogError:    []string{"Insert group error"},
		Origin:      "CreateAuthorWithBooks",
	}

	authorWithBooksUsecaseMock.On("CreateAuthorWithBooks", callExpected).Return(output, err)

	return authorWithBooksUsecaseMock, uuidMock, err
}
