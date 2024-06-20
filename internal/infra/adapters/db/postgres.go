package db

import (
	"context"
	"fmt"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/configs"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/logger"
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
			conf.DBHost, conf.DBPort, conf.DBUser, conf.DBPassword, conf.DBName)
		maxConns, _ := strconv.Atoi(conf.MaxConnections)
		minConns, _ := strconv.Atoi(conf.MinConnections)
		maxConnLifetime, _ := strconv.Atoi(conf.MaxConnLifetime)
		maxConnIdleTime, _ := strconv.Atoi(conf.MaxConnIdleTime)
		connConfig, err := pgxpool.ParseConfig(connStr)
		if err != nil {
			log.ErrorText(fmt.Sprintf("Error parsing connection string: %v", err))
			os.Exit(util.EXIT_FAILURE)
		}
		connConfig.MinConns = int32(minConns)
		connConfig.MaxConns = int32(maxConns)
		connConfig.MaxConnLifetime = time.Duration(maxConnLifetime) * time.Minute
		connConfig.MaxConnIdleTime = time.Duration(maxConnIdleTime) * time.Minute
		connConfig.HealthCheckPeriod = 10 * time.Minute
		connConfig.ConnConfig.RuntimeParams["application_name"] = "store-ms"
		pool, err = pgxpool.NewWithConfig(context.Background(), connConfig)
		if err != nil {
			log.ErrorText(fmt.Sprintf("Error creating connection pool: %v", err))
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
