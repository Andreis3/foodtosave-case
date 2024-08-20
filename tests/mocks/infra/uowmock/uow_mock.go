package uowmock

import (
	"github.com/stretchr/testify/mock"

	"github.com/andreis3/foodtosave-case/internal/domain/uow"
	"github.com/andreis3/foodtosave-case/internal/util"
)

type UnitOfWorkMock struct {
	mock.Mock
}

func (u *UnitOfWorkMock) Register(name string, callback uow.RepositoryFactory) {
	u.Called(name, callback)
}

func (u *UnitOfWorkMock) GetRepository(name string) any {
	args := u.Called(name)
	return args.Get(0).(any)
}

func (u *UnitOfWorkMock) Do(callback func(uow uow.IUnitOfWork) *util.ValidationError) *util.ValidationError {
	args := u.Called(callback)
	return args.Get(0).(*util.ValidationError)
}

func (u *UnitOfWorkMock) Rollback() *util.ValidationError {
	args := u.Called()
	return args.Get(0).(*util.ValidationError)
}

func (u *UnitOfWorkMock) CommitOrRollback() *util.ValidationError {
	args := u.Called()
	return args.Get(0).(*util.ValidationError)
}
