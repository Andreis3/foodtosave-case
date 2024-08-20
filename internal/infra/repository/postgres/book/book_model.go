package book

import (
	"time"

	"github.com/andreis3/foodtosave-case/internal/domain/entity"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type BookModel struct {
	ID        *string    `db:"id"`
	Title     *string    `db:"title"`
	Gender    *string    `db:"gender"`
	AuthorID  *string    `db:"author_id"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

func MapperBookModel(book entity.Book, authorId string) *BookModel {
	dateTime := util.FormatDateTime()
	return &BookModel{
		ID:        &book.ID,
		Title:     &book.Title,
		Gender:    &book.Gender,
		AuthorID:  &authorId,
		CreatedAt: &dateTime,
		UpdatedAt: &dateTime,
	}
}

func (b *BookModel) ToEntity() entity.Book {
	return entity.Book{
		ID:     *b.ID,
		Title:  *b.Title,
		Gender: *b.Gender,
	}
}
