package authorhandler

import "net/http"

type ICreateAuthorHandler interface {
	CreateAuthorWithBooks(w http.ResponseWriter, r *http.Request)
}

type IGetOneAuthorHandler interface {
	GetOneAuthorAllBooks(w http.ResponseWriter, r *http.Request)
}
