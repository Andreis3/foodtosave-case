//go:build unit
// +build unit

package query_test

import (
	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/domain/entity"
	"github.com/andreis3/foodtosave-case/internal/util"
	"github.com/andreis3/foodtosave-case/tests/mocks/domain/usecasemock"
)

func ContextGetSuccess() *usecasemock.GetOneAuthorAllBooksUseCaseMock {
	authorWithBooksUsecaseMock := new(usecasemock.GetOneAuthorAllBooksUseCaseMock)
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
		},
	}

	authorWithBooksUsecaseMock.On("GetOneAuthorAllBooks", "1").Return(output, (*util.ValidationError)(nil))

	return authorWithBooksUsecaseMock
}

func ContextGetError() (*usecasemock.GetOneAuthorAllBooksUseCaseMock, *util.ValidationError) {
	authorWithBooksUsecaseMock := new(usecasemock.GetOneAuthorAllBooksUseCaseMock)
	output := aggregate.AuthorBookAggregate{}

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
