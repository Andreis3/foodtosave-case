package mapper

import (
	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/domain/entity"
	"github.com/andreis3/foodtosave-case/internal/infra/dto"
)

func MapperDtoInputToAggregate(input dto.AuthorInput) aggregate.AuthorBookAggregate {
	var authorBookAggregate = aggregate.AuthorBookAggregate{}
	authorBookAggregate.Author = entity.Author{
		ID:          "",
		Name:        input.Name,
		Nationality: input.Nationality,
	}
	authorBookAggregate.Books = make([]entity.Book, len(input.Books))
	for i, book := range input.Books {
		authorBookAggregate.Books[i] = entity.Book{
			ID:     "",
			Title:  book.Title,
			Gender: book.Gender,
		}
	}
	return authorBookAggregate
}

func MapperAggregateToDtoOutput(agg aggregate.AuthorBookAggregate) dto.AuthorOutput {
	output := dto.AuthorOutput{
		ID:          agg.Author.ID,
		Name:        agg.Author.Name,
		Nationality: agg.Author.Nationality,
	}

	for _, book := range agg.Books {
		output.Books = append(output.Books, dto.BookOutput{
			ID:     book.ID,
			Title:  book.Title,
			Gender: book.Gender,
		})
	}
	return output
}
