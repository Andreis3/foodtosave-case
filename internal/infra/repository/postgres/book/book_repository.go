package book

import (
	"context"
	"fmt"
	"github.com/andreis3/foodtosave-case/internal/domain/entity"
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/db"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/observability"
	"github.com/andreis3/foodtosave-case/internal/util"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
)

type BookRepository struct {
	DB db.IInstructionDB
	*pgconn.PgError
	metrics observability.IMetricAdapter
}

func NewBookRepository(metrics observability.IMetricAdapter) *BookRepository {
	return &BookRepository{
		metrics: metrics,
	}
}
func (r *BookRepository) InsertBook(data entity.Book, authorId string) (*BookModel, *util.ValidationError) {
	start := time.Now()
	model := MapperBookModel(data, authorId)
	query := `INSERT INTO books (id, title, gender, author_id, created_at, updated_at) 
				VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`
	rows, _ := r.DB.Query(context.Background(), query,
		model.ID,
		model.Title,
		model.Gender,
		model.AuthorID,
		model.CreatedAt,
		model.UpdatedAt)
	defer rows.Close()
	group, err := pgx.CollectOneRow[BookModel](rows, pgx.RowToStructByName[BookModel])
	//ERROR: duplicate key value violates unique constraint "groups_name_code_key" (SQLSTATE 23505)
	if errors.As(err, &r.PgError) {
		return &BookModel{}, &util.ValidationError{
			Code:        fmt.Sprintf("PIDB-%s", r.Code),
			Origin:      "BookRepository.InsertBook",
			Status:      http.StatusInternalServerError,
			LogError:    []string{fmt.Sprintf("%s, %s", r.Message, r.Detail)},
			ClientError: []string{"Internal Server Error"},
		}
	}
	end := time.Now()
	duration := float64(end.Sub(start).Milliseconds())
	r.metrics.HistogramInstructionTableDuration(context.Background(), "postgres", "books", "insert", duration)
	return &group, nil
}
func (r *BookRepository) SelectAllBooksByAuthorID(authorId string) ([]BookModel, *util.ValidationError) {
	start := time.Now()
	query := `SELECT * FROM books WHERE author_id = $1`
	rows, _ := r.DB.Query(context.Background(), query, authorId)
	defer rows.Close()
	group, err := pgx.CollectOneRow[[]BookModel](rows, pgx.RowToStructByName[[]BookModel])
	if errors.As(err, &r.PgError) {
		return nil, &util.ValidationError{
			Code:        fmt.Sprintf("PIDB-%s", r.Code),
			Origin:      "BookRepository.SelectOneGroupByNameAndCode",
			Status:      http.StatusInternalServerError,
			LogError:    []string{fmt.Sprintf("%s, %s", r.Message, r.Detail)},
			ClientError: []string{"Internal Server Error"},
		}
	}
	end := time.Now()
	duration := float64(end.Sub(start).Milliseconds())
	r.metrics.HistogramInstructionTableDuration(context.Background(), "postgres", "books", "select", duration)
	return group, nil
}
