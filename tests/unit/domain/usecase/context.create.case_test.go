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

func ContextCreatedSuccess(authorRepoMock *authormock.AuthorRepositoryMock,
	bookRepoMock *bookmock.BookRepositoryMock,
	redisMock *cachemock.CacheMock,
	metrics *observabilitymock.PrometheusAdapterMock) *uowmock.UnitOfWorkMock {
	unitOfWork := new(uowmock.UnitOfWorkMock)

	authorRepoMock.On("InsertAuthor", mock.Anything).Return(&author.AuthorModel{}, (*util.ValidationError)(nil))
	bookRepoMock.On("InsertBook", mock.Anything, mock.Anything).Return(&book.BookModel{}, (*util.ValidationError)(nil))
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

func ContextReturnErrorAuthorRepositoryInsertAuthor(authorRepoMock *authormock.AuthorRepositoryMock,
	bookRepoMock *bookmock.BookRepositoryMock,
	redisMock *cachemock.CacheMock,
	metrics *observabilitymock.PrometheusAdapterMock) *uowmock.UnitOfWorkMock {
	unitOfWork := new(uowmock.UnitOfWorkMock)
	err := &util.ValidationError{
		Code:        "PIDB-235",
		Status:      500,
		LogError:    []string{"Insert author error"},
		ClientError: []string{"Internal Server Error"},
	}

	authorRepoMock.On("InsertAuthor", mock.Anything).Return(&author.AuthorModel{}, err)
	bookRepoMock.On("InsertBook", mock.Anything, mock.Anything).Return(&book.BookModel{}, (*util.ValidationError)(nil))
	redisMock.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	metrics.On("HistogramOperationDuration", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	unitOfWork.On("Do", mock.AnythingOfType("func(uow.IUnitOfWork) *util.ValidationError")).Return(err).Run(func(args mock.Arguments) {
		callback := args.Get(0).(func(uow.IUnitOfWork) *util.ValidationError)
		callback(unitOfWork)
	}).Once()

	unitOfWork.On("GetRepository", util.AUTH_REPOSITORY_KEY).Return(authorRepoMock)
	unitOfWork.On("GetRepository", util.BOOK_REPOSITORY_KEY).Return(bookRepoMock)

	return unitOfWork
}
