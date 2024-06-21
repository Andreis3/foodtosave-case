package dto

import "github.com/andreis3/foodtosave-case/internal/domain/aggregate"

type AuthorOutput struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
	Books       []struct {
		ID     string `json:"id"`
		Title  string `json:"title"`
		Gender string `json:"gender"`
	} `json:"books"`
}

func (a *AuthorOutput) MapperToAggregateAuthor(data aggregate.AuthorBookAggregate) AuthorOutput {
	output := AuthorOutput{
		ID:          data.Author.ID,
		Name:        data.Author.Name,
		Nationality: data.Author.Nationality,
	}
	for _, book := range data.Books {
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
