//go:build unit
// +build unit

package query_test

import (
	"github.com/andreis3/foodtosave-case/internal/app/query"
	"github.com/andreis3/foodtosave-case/internal/infra/dto"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("APP :: COMMAND :: GET_ONE_AUTHOR_ALL_BOOKS_USECASE", func() {
	Describe("#Execute", func() {
		Context("Should call the method GetOneAuthorAllBooks", func() {
			It("Should return a new author and books", func() {
				usecaseMock := ContextGetSuccess()
				query := query.NewGetOneAuthorAllBooksQuery(usecaseMock)

				outputExpected := dto.AuthorOutput{
					ID:          "1",
					Name:        "Author 1",
					Nationality: "Brazilian",
					Books: []dto.BookOutput{
						{
							ID:     "1",
							Title:  "Book 1",
							Gender: "Terror",
						},
					},
				}

				result, err := query.Execute("1")

				Expect(err).To(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(outputExpected))
			})

			It("Should return error", func() {
				usecaseMock, errContext := ContextGetError()
				query := query.NewGetOneAuthorAllBooksQuery(usecaseMock)

				//agg := aggregate.NewAuthorBookAggregate(uuidMock)
				//
				//input := dto.AuthorInput{
				//	Name:        "Author 1",
				//	Nationality: "Brazilian",
				//	Books: []dto.BookInput{
				//		{
				//			Title:  "Book 1",
				//			Gender: "Terror",
				//		},
				//		{
				//			Title:  "Book 2",
				//			Gender: "Comedy",
				//		},
				//	},
				//}
				//callExpected := agg.MapperDtoInputToAggregate(input)

				result, err := query.Execute("1")

				Expect(err).ToNot(BeNil())
				Expect(err).To(Equal(errContext))
				Expect(result).To(Equal(dto.AuthorOutput{}))
				Expect(usecaseMock.AssertNumberOfCalls(GinkgoT(), "GetOneAuthorAllBooks", 1)).To(BeTrue())
				Expect(usecaseMock.AssertCalled(GinkgoT(), "GetOneAuthorAllBooks", "1")).To(BeTrue())

			})
		})
	})
})
