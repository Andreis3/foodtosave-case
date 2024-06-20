package util

import "net/http"

type RouteType []RouteFields
type RouteFields struct {
	Method      string
	Path        string
	Controller  any
	Description string
	Type        string
	Middlewares []func(http.Handler) http.Handler
}
