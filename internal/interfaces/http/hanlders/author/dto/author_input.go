package dto

import (
	"github.com/andreis3/foodtosave-case/internal/domain/aggregate"
	"github.com/andreis3/foodtosave-case/internal/domain/entity"
)

type AuthorInput struct {
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
	Books       []struct {
		Title  string `json:"title"`
		Gender string `json:"gender"`
	} `json:"books"`
}

func (a *AuthorInput) MapperToAggregateAuthor() aggregate.AuthorBookAggregate {
	author := entity.Author{
		Name:        a.Name,
		Nationality: a.Nationality,
	}

	var books []entity.Book

	for _, book := range a.Books {
		books = append(books, entity.Book{
			Title:  book.Title,
			Gender: book.Gender,
		})
	}

	return aggregate.AuthorBookAggregate{
		Author: author,
		Books:  books,
	}
}
