//go:build unit
// +build unit

package usecase_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/domain/entity"
	"github.com/andreis3/foodtosave-case/internal/domain/usecase"
	"github.com/andreis3/foodtosave-case/internal/util"
	"github.com/andreis3/foodtosave-case/tests/mocks/infra/common/observabilitymock"
	"github.com/andreis3/foodtosave-case/tests/mocks/infra/repository/postgres/authormock"
	"github.com/andreis3/foodtosave-case/tests/mocks/infra/repository/postgres/bookmock"
	"github.com/andreis3/foodtosave-case/tests/mocks/infra/repository/redis/cachemock"
)

var _ = Describe("DOMAIN :: USECASE :: CREATE_AUTHOR_WITH_BOOKS_USECASE", func() {
	Context("When I call the method CreateAuthorWithBooks", func() {
		It("Should insert a new author with books not return errors", func() {
			authorRepositoryMock := new(authormock.AuthorRepositoryMock)
			bookRepositoryMock := new(bookmock.BookRepositoryMock)
			redisMock := new(cachemock.CacheMock)
			metrics := new(observabilitymock.PrometheusAdapterMock)

			unitOfWorkMock := ContextCreatedSuccess(authorRepositoryMock, bookRepositoryMock, redisMock, metrics)

			agg := aggregate.AuthorBookAggregate{
				Author: entity.Author{
					ID:          "",
					Name:        "Author 1",
					Nationality: "Brazilian",
				},
				Books: []entity.Book{
					{
						ID:     "",
						Title:  "Book 1",
						Gender: "Terror",
					},
				},
			}

			usecase := usecase.NewCreateAuthorWithBookUseCase(unitOfWorkMock, redisMock, metrics)

			out, err := usecase.CreateAuthorWithBooks(agg)

			Expect(err).To(BeNil())
			Expect(out).ToNot(BeNil())
			Expect(out.Author.ID).To(Equal("1"))
			Expect(out.Author.Name).To(Equal("Author 1"))
			Expect(out.Books[0].ID).To(Equal("1"))
			Expect(out.Books[0].Title).To(Equal("Book 1"))
			Expect(out.Books[0].Gender).To(Equal("Terror"))

		})

		It("Should return an error when the method CreateAuthorWithBooks is call", func() {
			authorRepositoryMock := new(authormock.AuthorRepositoryMock)
			bookRepositoryMock := new(bookmock.BookRepositoryMock)
			redisMock := new(cachemock.CacheMock)
			metrics := new(observabilitymock.PrometheusAdapterMock)

			unitOfWorkMock := ContextReturnErrorAuthorRepositoryInsertAuthor(authorRepositoryMock, bookRepositoryMock, redisMock, metrics)

			agg := aggregate.AuthorBookAggregate{
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

			usecase := usecase.NewCreateAuthorWithBookUseCase(unitOfWorkMock, redisMock, metrics)

			out, err := usecase.CreateAuthorWithBooks(agg)

			expectedError := &util.ValidationError{
				Code:        "PIDB-235",
				Status:      500,
				LogError:    []string{"Insert author error"},
				ClientError: []string{"Internal Server Error"},
			}

			Expect(err).ToNot(BeNil())
			Expect(out).To(BeZero())
			Expect(err).To(Equal(expectedError))
			Expect(redisMock.AssertNotCalled(GinkgoT(), "Set", "author:1", out)).To(BeTrue())
		})
	})
})
