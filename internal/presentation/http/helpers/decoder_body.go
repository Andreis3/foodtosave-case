package helpers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/andreis3/foodtosave-case/internal/util"
)

func DecoderBodyRequest[T any](req *http.Request) (T, *util.ValidationError) {
	defer req.Body.Close()
	var result T
	var jsonUnmarshalTypeError *json.UnmarshalTypeError
	var jsonSyntaxError *json.SyntaxError
	err := json.NewDecoder(req.Body).Decode(&result)
	switch {
	case errors.As(err, &jsonSyntaxError):
		return result, &util.ValidationError{
			Code:        "DJ-400",
			Origin:      "DecoderBodyRequest",
			Status:      http.StatusBadRequest,
			ClientError: []string{"invalid json syntax"},
			LogError:    []string{jsonSyntaxError.Error()},
		}
	case errors.As(err, &jsonUnmarshalTypeError):
		return result, &util.ValidationError{
			Code:        "DJ-401",
			Origin:      "DecoderBodyRequest",
			Status:      http.StatusBadRequest,
			ClientError: []string{"invalid json field type"},
			LogError:    []string{jsonUnmarshalTypeError.Error()},
		}
	case err != nil:
		return result, &util.ValidationError{
			Code:        "DJ-402",
			Origin:      "DecoderBodyRequest",
			Status:      http.StatusBadRequest,
			LogError:    []string{err.Error()},
			ClientError: []string{"invalid json"},
		}
	}
	return result, nil
}
