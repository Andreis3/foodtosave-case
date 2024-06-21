package handler_author

import "net/http"

type ICreateAuthorHandler interface {
	CreateAuthor(w http.ResponseWriter, r *http.Request)
}

type IGetOneAuthorHandler interface {
	GetOneAuthor(w http.ResponseWriter, r *http.Request)
}
