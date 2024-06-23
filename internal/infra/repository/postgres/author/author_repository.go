package author

import (
	"context"
	"errors"
	"fmt"
	"github.com/andreis3/foodtosave-case/internal/domain/entity"
	"github.com/andreis3/foodtosave-case/internal/infra/adapters/db"
	"github.com/andreis3/foodtosave-case/internal/infra/commons/observability"
	"github.com/andreis3/foodtosave-case/internal/util"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type AuthorRepository struct {
	DB *db.Queries
	*pgconn.PgError
	metrics observability.IMetricAdapter
}

func NewAuthorRepository(metrics observability.IMetricAdapter) *AuthorRepository {
	return &AuthorRepository{
		metrics: metrics,
	}
}
func (r *AuthorRepository) InsertAuthor(data entity.Author) (*AuthorModel, *util.ValidationError) {
	start := time.Now()
	model := MapperAuthorModel(data)
	query := `INSERT INTO authors (id, name, nationality, created_at, updated_at) 
				VALUES ($1, $2, $3, $4, $5) RETURNING *`
	rows, _ := r.DB.Query(context.Background(), query,
		model.ID,
		model.Name,
		model.Nationality,
		model.CreatedAt,
		model.UpdatedAt)
	defer rows.Close()
	group, err := pgx.CollectOneRow[AuthorModel](rows, pgx.RowToStructByName[AuthorModel])
	//ERROR: duplicate key value violates unique constraint "groups_name_code_key" (SQLSTATE 23505)
	if errors.As(err, &r.PgError) {
		return &AuthorModel{}, &util.ValidationError{
			Code:        fmt.Sprintf("PIDB-%s", r.Code),
			Origin:      "AuthorRepository.InsertAuthor",
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
			ClientError: []string{"Internal Server Error"},
		}
	}
	end := time.Now()
	duration := float64(end.Sub(start).Milliseconds())
	r.metrics.HistogramInstructionTableDuration(context.Background(), "postgres", "books", "select", duration)
	return &author, nil
}
