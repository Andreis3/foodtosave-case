//go:build unit
// +build unit

package usecase_test

import (
	"github.com/andreis3/foodtosave-case/internal/domain/usecase"
	"github.com/andreis3/foodtosave-case/tests/mocks/infra/common/observabilitymock"
	"github.com/andreis3/foodtosave-case/tests/mocks/infra/repository/postgres/authormock"
	"github.com/andreis3/foodtosave-case/tests/mocks/infra/repository/postgres/bookmock"
	"github.com/andreis3/foodtosave-case/tests/mocks/infra/repository/redis/cachemock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("APP :: USECASE :: GET_ONE_AUTHOR_ALL_BOOKS_USECASE", func() {
	Describe("#Execute", func() {
		Context("When I call the method GetOneAuthorAllBooks of the get_one_author_all_books_service", func() {
			It("Should search a new author not return errors", func() {
				authorRepositoryMock := new(authormock.AuthorRepositoryMock)
				bookRepositoryMock := new(bookmock.BookRepositoryMock)
				redisMock := new(cachemock.CacheMock)
				metrics := new(observabilitymock.PrometheusAdapterMock)

				unitOfWorkMock := ContextGetCacheSuccess(authorRepositoryMock, bookRepositoryMock, redisMock, metrics)

				authorID := "1"

				usecase := usecase.NewGetOneAuthorAllBooksUsecase(unitOfWorkMock, redisMock, metrics)

				out, err := usecase.GetOneAuthorAllBooks(authorID)

				Expect(err).To(BeNil())
				Expect(out).ToNot(BeNil())
				Expect(out.ID).To(Equal("1"))
				Expect(out.Name).To(Equal("Author 1"))
				Expect(out.Nationality).To(Equal("Brazilian"))
				Expect(out.Books[0].ID).To(Equal("1"))
				Expect(out.Books[0].Title).To(Equal("Book 1"))
				Expect(out.Books[0].Gender).To(Equal("Terror"))
			})

			It("Should return data from postgres when the cache is empty", func() {
				authorRepositoryMock := new(authormock.AuthorRepositoryMock)
				bookRepositoryMock := new(bookmock.BookRepositoryMock)
				redisMock := new(cachemock.CacheMock)
				metrics := new(observabilitymock.PrometheusAdapterMock)

				unitOfWorkMock := ContextGetPostgresSuccess(authorRepositoryMock, bookRepositoryMock, redisMock, metrics)

				authorID := "1"

				usecase := usecase.NewGetOneAuthorAllBooksUsecase(unitOfWorkMock, redisMock, metrics)

				out, err := usecase.GetOneAuthorAllBooks(authorID)

				Expect(err).To(BeNil())
				Expect(out).ToNot(BeNil())

				Expect(out.ID).To(Equal("1"))
				Expect(out.Name).To(Equal("Author 1"))
				Expect(out.Nationality).To(Equal("Brazilian"))
				Expect(out.Books[0].ID).To(Equal("1"))
				Expect(out.Books[0].Title).To(Equal("Book 1"))
				Expect(out.Books[0].ID).To(Equal("1"))
				Expect(out.Books[0].Title).To(Equal("Book 1"))
				Expect(out.Books[0].Gender).To(Equal("Terror"))
			})

			It("Should return an error when the method GetOneAuthorAllBooks is call", func() {
				authorRepositoryMock := new(authormock.AuthorRepositoryMock)
				bookRepositoryMock := new(bookmock.BookRepositoryMock)
				redisMock := new(cachemock.CacheMock)
				metrics := new(observabilitymock.PrometheusAdapterMock)

				unitOfWorkMock, errExpected := ContextReturnErrorAuthorRepositorySelectOneAuthorByID(authorRepositoryMock, bookRepositoryMock, redisMock, metrics)

				authorID := "1"

				usecase := usecase.NewGetOneAuthorAllBooksUsecase(unitOfWorkMock, redisMock, metrics)

				out, err := usecase.GetOneAuthorAllBooks(authorID)

				Expect(err).ToNot(BeNil())
				Expect(out).To(BeZero())
				Expect(err).To(Equal(errExpected))
			})
		})
	})
})
