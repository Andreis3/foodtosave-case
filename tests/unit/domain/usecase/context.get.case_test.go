//go:build unit
// +build unit

package usecase_test

import (
	"github.com/andreis3/foodtosave-case/internal/domain/uow"
	"github.com/andreis3/foodtosave-case/internal/infra/repository/postgres/author"
	"github.com/andreis3/foodtosave-case/internal/infra/repository/postgres/book"
	"github.com/andreis3/foodtosave-case/internal/util"
	"github.com/andreis3/foodtosave-case/tests/mocks/infra/common/observabilitymock"
	"github.com/andreis3/foodtosave-case/tests/mocks/infra/repository/postgres/authormock"
	"github.com/andreis3/foodtosave-case/tests/mocks/infra/repository/postgres/bookmock"
	"github.com/andreis3/foodtosave-case/tests/mocks/infra/repository/redis/cachemock"
	"github.com/andreis3/foodtosave-case/tests/mocks/infra/uowmock"
	"github.com/stretchr/testify/mock"
)

func ContextGetCacheSuccess(authorRepoMock *authormock.AuthorRepositoryMock,
	bookRepoMock *bookmock.BookRepositoryMock,
	redisMock *cachemock.CacheMock,
	metrics *observabilitymock.PrometheusAdapterMock) *uowmock.UnitOfWorkMock {
	unitOfWork := new(uowmock.UnitOfWorkMock)

	cache := `{"Author":{"ID":"1","Name":"Author 1","Nationality":"test 1"},"Books":[{"ID":"1","Title":"test 1","Gender":"test 1"}]}`

	authorRepoMock.On("SelectOneAuthorByID", mock.Anything).Return(&author.AuthorModel{}, (*util.ValidationError)(nil))
	bookRepoMock.On("SelectAllBooksByAuthorID", mock.Anything).Return([]*book.BookModel{}, (*util.ValidationError)(nil))
	redisMock.On("Get", mock.Anything).Return(cache, nil)
	metrics.On("HistogramOperationDuration", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	unitOfWork.On("Do", mock.AnythingOfType("func(uow.IUnitOfWork) *util.ValidationError")).Return((*util.ValidationError)(nil)).Run(func(args mock.Arguments) {
		callback := args.Get(0).(func(uow.IUnitOfWork) *util.ValidationError)
		callback(unitOfWork)
	}).Once()

	unitOfWork.On("GetRepository", util.AUTH_REPOSITORY_KEY).Return(authorRepoMock)
	unitOfWork.On("GetRepository", util.BOOK_REPOSITORY_KEY).Return(bookRepoMock)
	return unitOfWork
}

func ContextGetPostgresSuccess(authorRepoMock *authormock.AuthorRepositoryMock,
	bookRepoMock *bookmock.BookRepositoryMock,
	redisMock *cachemock.CacheMock,
	metrics *observabilitymock.PrometheusAdapterMock) *uowmock.UnitOfWorkMock {
	unitOfWork := new(uowmock.UnitOfWorkMock)

	id := "1"
	name := "Author 1"
	nationality := "Brazilian"
	authorModel := &author.AuthorModel{
		ID:          &id,
		Name:        &name,
		Nationality: &nationality,
	}

	title := "Book 1"
	gender := "Terror"

	bookModel := []book.BookModel{
		{
			ID:     &id,
			Title:  &title,
			Gender: &gender,
		},
	}

	authorRepoMock.On("SelectOneAuthorByID", mock.Anything).Return(authorModel, (*util.ValidationError)(nil))
	bookRepoMock.On("SelectAllBooksByAuthorID", mock.Anything).Return(bookModel, (*util.ValidationError)(nil))
	redisMock.On("Get", mock.Anything).Return("", nil)
	redisMock.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	metrics.On("HistogramOperationDuration", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	unitOfWork.On("Do", mock.AnythingOfType("func(uow.IUnitOfWork) *util.ValidationError")).Return((*util.ValidationError)(nil)).Run(func(args mock.Arguments) {
		callback := args.Get(0).(func(uow.IUnitOfWork) *util.ValidationError)
		callback(unitOfWork)
	}).Once()

	unitOfWork.On("GetRepository", util.AUTH_REPOSITORY_KEY).Return(authorRepoMock)
	unitOfWork.On("GetRepository", util.BOOK_REPOSITORY_KEY).Return(bookRepoMock)

	return unitOfWork
}

func ContextReturnErrorAuthorRepositorySelectOneAuthorByID(authorRepoMock *authormock.AuthorRepositoryMock,
	bookRepoMock *bookmock.BookRepositoryMock,
	redisMock *cachemock.CacheMock,
	metrics *observabilitymock.PrometheusAdapterMock) (*uowmock.UnitOfWorkMock, *util.ValidationError) {
	unitOfWork := new(uowmock.UnitOfWorkMock)

	err := &util.ValidationError{
		Code:        "PIDB-235",
		Status:      500,
		LogError:    []string{"Error select author by id"},
		ClientError: []string{"Internal Server Error"},
	}

	authorRepoMock.On("SelectOneAuthorByID", mock.Anything).Return(&author.AuthorModel{}, err)
	bookRepoMock.On("SelectAllBooksByAuthorID", mock.Anything).Return([]*book.BookModel{}, (*util.ValidationError)(nil))
	redisMock.On("Get", mock.Anything).Return("", nil)
	metrics.On("HistogramOperationDuration", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	unitOfWork.On("Do", mock.AnythingOfType("func(uow.IUnitOfWork) *util.ValidationError")).Return(err).Run(func(args mock.Arguments) {
		callback := args.Get(0).(func(uow.IUnitOfWork) *util.ValidationError)
		callback(unitOfWork)
	}).Once()

	unitOfWork.On("GetRepository", util.AUTH_REPOSITORY_KEY).Return(authorRepoMock)
	unitOfWork.On("GetRepository", util.BOOK_REPOSITORY_KEY).Return(bookRepoMock)

	return unitOfWork, err
}
