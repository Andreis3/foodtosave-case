package postgres

import (
	"context"
	"fmt"
	"github.com/andreis3/foodtosave-case/internal/infra/common/configs"
	"github.com/andreis3/foodtosave-case/internal/infra/common/logger"
	"github.com/andreis3/foodtosave-case/internal/util"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	pool *pgxpool.Pool
}

var singleton sync.Once
var pool *pgxpool.Pool

func NewPostgresDB(conf configs.Conf) *Postgres {
	log := logger.NewLogger()
	singleton.Do(func() {
		connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			conf.PostgresHost, conf.PostgresPort, conf.PostgresUser, conf.PostgresPassword, conf.PostgresDBName)
		maxConns, _ := strconv.Atoi(conf.PostgresMaxConnections)
		minConns, _ := strconv.Atoi(conf.PostgresMinConnections)
		maxConnLifetime, _ := strconv.Atoi(conf.PostgresMaxConnLifetime)
		maxConnIdleTime, _ := strconv.Atoi(conf.PostgresMaxConnIdleTime)
		connConfig, err := pgxpool.ParseConfig(connStr)
		if err != nil {
			log.ErrorText(fmt.Sprintf("NotificationErrors parsing connection string: %v", err))
			os.Exit(util.EXIT_FAILURE)
		}
		connConfig.MinConns = int32(minConns)
		connConfig.MaxConns = int32(maxConns)
		connConfig.MaxConnLifetime = time.Duration(maxConnLifetime) * time.Minute
		connConfig.MaxConnIdleTime = time.Duration(maxConnIdleTime) * time.Minute
		connConfig.HealthCheckPeriod = 10 * time.Minute
		connConfig.ConnConfig.RuntimeParams["application_name"] = "foodtosave-case"
		pool, err = pgxpool.NewWithConfig(context.Background(), connConfig)
		if err != nil {
			log.ErrorText(fmt.Sprintf("NotificationErrors creating connection pool: %v", err))
			os.Exit(util.EXIT_FAILURE)
		}
	})
	return &Postgres{pool: pool}
}

func (p *Postgres) InstanceDB() any {
	return p.pool
}

func (p *Postgres) Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error) {
	return p.pool.Exec(ctx, sql, arguments...)
}
func (p *Postgres) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return p.pool.Query(ctx, sql, args...)
}
func (p *Postgres) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return p.pool.QueryRow(ctx, sql, args...)
}
func (p *Postgres) Close() {
	p.pool.Close()
}
