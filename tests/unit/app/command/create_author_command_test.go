//go:build unit
// +build unit

package command_test

import (
	"github.com/andreis3/foodtosave-case/internal/app/command"
	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/infra/dto"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("APP :: COMMAND :: CREATE_AUTHOR_WITH_BOOKS_COMMAND", func() {
	Describe("#Execute", func() {
		Context("Should call the method CreateAuthorWithBooks", func() {
			It("Should return a new author and books", func() {
				usecaseMock, uuidMock, output := ContextCreateSuccess()
				command := command.NewCreateAuthorWithBooksCommand(usecaseMock, uuidMock)

				input := dto.AuthorInput{
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
				}

				result, err := command.Execute(input)

				Expect(err).To(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(output))
			})

			It("Should return error", func() {
				usecaseMock, uuidMock, errContext := ContextCreateError()
				command := command.NewCreateAuthorWithBooksCommand(usecaseMock, uuidMock)

				agg := aggregate.NewAuthorBookAggregate(uuidMock)

				input := dto.AuthorInput{
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
				}
				callExpected := agg.MapperDtoInputToAggregate(input)

				result, err := command.Execute(input)

				Expect(err).ToNot(BeNil())
				Expect(err).To(Equal(errContext))
				Expect(result).To(Equal(dto.AuthorOutput{}))
				Expect(usecaseMock.AssertNumberOfCalls(GinkgoT(), "CreateAuthorWithBooks", 1)).To(BeTrue())
				Expect(usecaseMock.AssertCalled(GinkgoT(), "CreateAuthorWithBooks", callExpected)).To(BeTrue())

			})
		})
	})
})
