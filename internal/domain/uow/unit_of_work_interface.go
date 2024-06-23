package uow

import "github.com/andreis3/foodtosave-case/internal/util"

type RepositoryFactory func(tx any) any

type IUnitOfWork interface {
	Register(name string, callback RepositoryFactory)
	GetRepository(name string) any
	Do(callback func(uow IUnitOfWork) *util.ValidationError) *util.ValidationError
	CommitOrRollback() *util.ValidationError
	Rollback() *util.ValidationError
}
