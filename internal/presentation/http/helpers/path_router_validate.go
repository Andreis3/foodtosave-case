package helpers

import (
	"net/http"

	"github.com/andreis3/foodtosave-case/internal/util"

	"github.com/google/uuid"
)

const (
	ID = "id"
)

func PathRouterValidate(r *http.Request, kind string) *util.ValidationError {
	switch kind {
	case ID:
		id := r.PathValue("id")
		if err := uuid.Validate(id); err != nil {
			return &util.ValidationError{
				Code:        "PR-001",
				Origin:      "PathRouterValidate",
				ClientError: []string{"invalid parameter id"},
				LogError:    []string{"invalid path parameter id"},
				Status:      http.StatusBadRequest,
			}
		}
	}
	return nil
}
