package helpers

import "net/http"

// Router types
const (
	HANDLER      = "handler"
	HANDLER_FUNC = "handlerFunc"
)

type RouteType []RouteFields
type RouteFields struct {
	Method      string
	Path        string
	Controller  any
	Description string
	Type        string
	Middlewares []func(http.Handler) http.Handler
}
