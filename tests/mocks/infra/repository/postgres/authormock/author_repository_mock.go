package authormock

import (
	"github.com/andreis3/foodtosave-case/internal/domain/entity"
	"github.com/andreis3/foodtosave-case/internal/infra/repository/postgres/author"
	"github.com/andreis3/foodtosave-case/internal/util"
	"github.com/stretchr/testify/mock"
)

type AuthorRepositoryMock struct {
	mock.Mock
}

func (r *AuthorRepositoryMock) InsertAuthor(data entity.Author) (string, *util.ValidationError) {
	args := r.Called(data)
	return args.Get(0).(string), args.Get(1).(*util.ValidationError)
}
func (r *AuthorRepositoryMock) SelectOneAuthorByID(authorId string) (*author.AuthorModel, *util.ValidationError) {
	args := r.Called(authorId)
	return args.Get(0).(*author.AuthorModel), args.Get(1).(*util.ValidationError)
}
