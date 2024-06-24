package book

import (
	"context"
	"errors"
	"fmt"
	"github.com/andreis3/foodtosave-case/internal/domain/entity"
	"github.com/andreis3/foodtosave-case/internal/domain/observability"
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/db/postgres"
	"github.com/andreis3/foodtosave-case/internal/util"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
)

type BookRepository struct {
	DB *postgres.Queries
	*pgconn.PgError
	metrics observability.IMetricAdapter
}

func NewBookRepository(metrics observability.IMetricAdapter) *BookRepository {
	return &BookRepository{
		metrics: metrics,
	}
}
func (r *BookRepository) InsertBook(data entity.Book, authorId string) (string, *util.ValidationError) {
	start := time.Now()
	model := MapperBookModel(data, authorId)
	var bookId string
	query := `INSERT INTO books (title, gender, author_id, created_at, updated_at) 
				VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.DB.QueryRow(context.Background(), query,
		model.Title,
		model.Gender,
		model.AuthorID,
		model.CreatedAt,
		model.UpdatedAt).Scan(&bookId)
	if errors.As(err, &r.PgError) {
		return "", &util.ValidationError{
			Code:        fmt.Sprintf("PIDB-%s", r.Code),
			Origin:      "BookRepository.InsertBook",
			Status:      http.StatusInternalServerError,
			LogError:    []string{fmt.Sprintf("%s, %s", r.Message, r.Detail)},
			ClientError: []string{"Internal Server NotificationErrors"},
		}
	}
	end := time.Now()
	duration := float64(end.Sub(start).Milliseconds())
	r.metrics.HistogramInstructionTableDuration(context.Background(), "postgres", "books", "insert", duration)
	return bookId, nil
}
func (r *BookRepository) SelectAllBooksByAuthorID(authorId string) ([]BookModel, *util.ValidationError) {
	start := time.Now()
	query := `SELECT id, title, gender FROM books WHERE author_id = $1`
	var books []BookModel
	rows, err := r.DB.Query(context.Background(), query, authorId)
	defer rows.Close()
	if errors.As(err, &r.PgError) {
		return nil, &util.ValidationError{
			Code:        fmt.Sprintf("PIDB-%s", r.Code),
			Origin:      "BookRepository.SelectOneGroupByNameAndCode",
			Status:      http.StatusInternalServerError,
			LogError:    []string{fmt.Sprintf("%s, %s", r.Message, r.Detail)},
			ClientError: []string{"Internal Server NotificationErrors"},
		}
	}
	defer rows.Close()
	for rows.Next() {
		book := BookModel{}
		err = rows.Scan(&book.ID, &book.Title, &book.Gender)
		books = append(books, book)
		if err != nil {
			return nil, &util.ValidationError{
				Code:        "PIDB-0001",
				Origin:      "BookRepository.SelectAllBooksByAuthorID",
				Status:      http.StatusInternalServerError,
				LogError:    []string{err.Error()},
				ClientError: []string{"Internal Server NotificationErrors"},
			}

		}
	}
	end := time.Now()
	duration := float64(end.Sub(start).Milliseconds())
	r.metrics.HistogramInstructionTableDuration(context.Background(), "postgres", "books", "select", duration)
	return books, nil
}
