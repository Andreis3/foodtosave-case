package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type IInstructionDB interface {
	Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error)
}
type Queries struct {
	IInstructionDB
}

func New(db IInstructionDB) *Queries {
	return &Queries{db}
}
