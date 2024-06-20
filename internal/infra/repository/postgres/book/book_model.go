package book

import (
	"github.com/andreis3/foodtosave-case/internal/domain/entity"
	"github.com/andreis3/foodtosave-case/internal/util"
	"time"
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
