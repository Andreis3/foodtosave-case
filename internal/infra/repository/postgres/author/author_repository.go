package author

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/andreis3/foodtosave-case/internal/domain/entity"
	"github.com/andreis3/foodtosave-case/internal/domain/observability"
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/db/postgres"
	"github.com/andreis3/foodtosave-case/internal/util"

	"github.com/jackc/pgx/v5/pgconn"
)

type AuthorRepository struct {
	DB *postgres.Queries
	*pgconn.PgError
	metrics observability.IMetricAdapter
}

func NewAuthorRepository(metrics observability.IMetricAdapter) *AuthorRepository {
	return &AuthorRepository{
		metrics: metrics,
	}
}
func (r *AuthorRepository) InsertAuthor(data entity.Author) (string, *util.ValidationError) {
	start := time.Now()
	model := MapperAuthorModel(data)
	var authorId string
	query := `INSERT INTO authors (name, nationality, created_at, updated_at) 
				VALUES ($1, $2, $3, $4 ) RETURNING id`
	err := r.DB.QueryRow(context.Background(), query,
		model.Name,
		model.Nationality,
		model.CreatedAt,
		model.UpdatedAt).Scan(&authorId)

	if errors.As(err, &r.PgError) {
		return "", &util.ValidationError{
			Code:        fmt.Sprintf("PIDB-%s", r.Code),
			Origin:      "AuthorRepository.InsertAuthor",
			Status:      http.StatusInternalServerError,
			LogError:    []string{fmt.Sprintf("%s, %s", r.Message, r.Detail)},
			ClientError: []string{"Internal Server"},
		}
	}
	end := time.Now()
	duration := float64(end.Sub(start).Milliseconds())
	r.metrics.HistogramInstructionTableDuration(context.Background(), "postgres", "author", "insert", duration)
	return authorId, nil
}
func (r *AuthorRepository) SelectOneAuthorByID(authorId string) (*AuthorModel, *util.ValidationError) {
	start := time.Now()
	var author AuthorModel
	query := `SELECT id, name, nationality FROM authors WHERE id = $1`
	rows, err := r.DB.Query(context.Background(), query, authorId)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&author.ID, &author.Name, &author.Nationality)
	}
	if errors.As(err, &r.PgError) {
		return nil, &util.ValidationError{
			Code:        fmt.Sprintf("PIDB-%s", r.Code),
			Origin:      "AuthorRepository.SelectOneAuthorByID",
			Status:      http.StatusInternalServerError,
			LogError:    []string{fmt.Sprintf("%s, %s", r.Message, r.Detail)},
			ClientError: []string{"Internal Server"},
		}
	}
	end := time.Now()
	duration := float64(end.Sub(start).Milliseconds())
	r.metrics.HistogramInstructionTableDuration(context.Background(), "postgres", "author", "select", duration)
	return &author, nil
}
