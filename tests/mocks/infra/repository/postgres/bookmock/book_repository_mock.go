package bookmock

import (
	"github.com/stretchr/testify/mock"

	"github.com/andreis3/foodtosave-case/internal/domain/entity"
	"github.com/andreis3/foodtosave-case/internal/infra/repository/postgres/book"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type BookRepositoryMock struct {
	mock.Mock
}

func (r *BookRepositoryMock) InsertBook(data entity.Book, authorId string) (string, *util.ValidationError) {
	args := r.Called(data, authorId)
	return args.Get(0).(string), args.Get(1).(*util.ValidationError)
}
func (r *BookRepositoryMock) SelectAllBooksByAuthorID(authorId string) ([]book.BookModel, *util.ValidationError) {
	args := r.Called(authorId)
	return args.Get(0).([]book.BookModel), args.Get(1).(*util.ValidationError)
}
