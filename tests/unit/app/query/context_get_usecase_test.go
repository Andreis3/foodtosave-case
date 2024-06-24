//go:build unit
// +build unit

package query_test

import (
	"github.com/andreis3/foodtosave-case/internal/infra/dto"
	"github.com/andreis3/foodtosave-case/internal/util"
	"github.com/andreis3/foodtosave-case/tests/mocks/domain/usecasemock"
)

func ContextGetSuccess() *usecasemock.GetOneAuthorAllBooksUsecaseMock {
	authorWithBooksUsecaseMock := new(usecasemock.GetOneAuthorAllBooksUsecaseMock)
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

	authorWithBooksUsecaseMock.On("GetOneAuthorAllBooks", "1").Return(output, (*util.ValidationError)(nil))

	return authorWithBooksUsecaseMock
}

func ContextGetError() (*usecasemock.GetOneAuthorAllBooksUsecaseMock, *util.ValidationError) {
	authorWithBooksUsecaseMock := new(usecasemock.GetOneAuthorAllBooksUsecaseMock)
	output := dto.AuthorOutput{}

	err := &util.ValidationError{
		Code:        "PIDB-235",
		Status:      500,
		ClientError: []string{"Internal Server Error"},
		LogError:    []string{"Insert group error"},
		Origin:      "CreateAuthorWithBooks",
	}

	authorWithBooksUsecaseMock.On("GetOneAuthorAllBooks", "1").Return(output, err)

	return authorWithBooksUsecaseMock, err
}
