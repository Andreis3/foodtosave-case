package repoauthor

import (
	"github.com/andreis3/foodtosave-case/internal/domain/entity"
	"github.com/andreis3/foodtosave-case/internal/util"
	"time"
)

type AuthorModel struct {
	ID          *string    `db:"id"`
	Name        *string    `db:"name"`
	Nationality *string    `db:"nationality"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}

func MapperAuthorModel(author entity.Author) *AuthorModel {
	dateTime := util.FormatDateTime()
	return &AuthorModel{
		ID:          &author.ID,
		Name:        &author.Name,
		Nationality: &author.Nationality,
		CreatedAt:   &dateTime,
		UpdatedAt:   &dateTime,
	}
}
