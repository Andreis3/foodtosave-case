package middleware

import (
	"fmt"
	"net/http"
)

func ValidatePath(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Validating path")
		next.ServeHTTP(w, r)
	})
}
