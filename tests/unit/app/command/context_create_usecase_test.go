//go:build unit
// +build unit

package command_test

import (
	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/infra/dto"
	"github.com/andreis3/foodtosave-case/internal/util"
	"github.com/andreis3/foodtosave-case/tests/mocks/domain/usecasemock"
	"github.com/andreis3/foodtosave-case/tests/mocks/infra/common/uuidmock"
	"github.com/stretchr/testify/mock"
)

func ContextCreateSuccess() (*usecasemock.CreateAuthorWithBooksUsecaseMock, *uuidmock.UUIDMock, dto.AuthorOutput) {
	authorWithBooksUsecaseMock := new(usecasemock.CreateAuthorWithBooksUsecaseMock)
	uuidMock := new(uuidmock.UUIDMock)
	output := dto.AuthorOutput{
		ID:          "1",
		Name:        "Author 1",
		Nationality: "Brazilian",
		Books: []dto.BookOutput{
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
	output := dto.AuthorOutput{}

	agg := aggregate.NewAuthorBookAggregate(uuidMock)
	uuidMock.On("Generate").Return("111")

	callExpected := agg.MapperDtoInputToAggregate(dto.AuthorInput{
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
