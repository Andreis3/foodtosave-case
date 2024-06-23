package aggregate

import (
	"github.com/andreis3/foodtosave-case/internal/domain/entity"
	"github.com/andreis3/foodtosave-case/internal/domain/notification"
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/uuid"
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/authorhandler/dto"
)

type AuthorBookAggregate struct {
	Author        entity.Author
	Books         []entity.Book
	uuidGenerator uuid.IUUID
}

func NewAuthorBookAggregate(uuidGenerator uuid.IUUID) *AuthorBookAggregate {
	return &AuthorBookAggregate{
		Author:        entity.Author{},
		Books:         []entity.Book{},
		uuidGenerator: uuidGenerator,
	}
}
func (a *AuthorBookAggregate) AddAuthorAndBooks(author entity.Author, books []entity.Book) {
	a.Author = author
	a.Books = books
}
func (a *AuthorBookAggregate) MapperDtoInputToAggregate(input dto.AuthorInput) AuthorBookAggregate {
	a.Author = entity.Author{
		ID:          a.uuidGenerator.Generate(),
		Name:        input.Name,
		Nationality: input.Nationality,
	}
	a.Books = make([]entity.Book, len(input.Books))
	for i, book := range input.Books {
		a.Books[i] = entity.Book{
			ID:     a.uuidGenerator.Generate(),
			Title:  book.Title,
			Gender: book.Gender,
		}
	}
	return *a
}
func (a *AuthorBookAggregate) MapperToDtoOutput() dto.AuthorOutput {
	output := dto.AuthorOutput{
		ID:          a.Author.ID,
		Name:        a.Author.Name,
		Nationality: a.Author.Nationality,
	}

	for _, book := range a.Books {
		output.Books = append(output.Books, struct {
			ID     string `json:"id"`
			Title  string `json:"title"`
			Gender string `json:"gender"`
		}{
			ID:     book.ID,
			Title:  book.Title,
			Gender: book.Gender,
		})
	}
	return output
}
func (a *AuthorBookAggregate) Validate() *notification.Error {
	authorValidate := a.Author.Validate()
	if len(a.Books) == 0 {
		authorValidate.AddErrors("books: minimum 1 book is required")
	}
	for index := range a.Books {
		authorValidate.MergeErrors(index, "books", a.Books[index].Validate())
	}
	return authorValidate
}
