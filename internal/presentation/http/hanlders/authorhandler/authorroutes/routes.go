package authorroutes

import (
	"github.com/andreis3/foodtosave-case/internal/presentation/http/hanlders/authorhandler"
	"github.com/andreis3/foodtosave-case/internal/presentation/http/helpers"
	"github.com/andreis3/foodtosave-case/internal/presentation/http/middleware"
	"net/http"
)

type Routes struct {
	createAuthorHandler authorhandler.CreateAuthorWithBooksHandler
	getOneAuthorHandler authorhandler.GetOneAuthorAllBooksHandler
}

func NewAuthorRoutes(
	createAuthorHandler authorhandler.CreateAuthorWithBooksHandler,
	getOneAuthorHandler authorhandler.GetOneAuthorAllBooksHandler) *Routes {
	return &Routes{
		createAuthorHandler: createAuthorHandler,
		getOneAuthorHandler: getOneAuthorHandler,
	}
}

func (r *Routes) AuthorRoutes() helpers.RouteType {
	return helpers.RouteType{
		{
			Method:      http.MethodPost,
			Path:        helpers.CREATE_AUTHOR_V1,
			Controller:  r.createAuthorHandler.CreateAuthorWithBooks,
			Description: "Create Author with Books",
			Type:        helpers.HANDLER_FUNC,
			Middlewares: []func(http.Handler) http.Handler{
				middleware.ValidatePath,
			},
		},
		{
			Method:      http.MethodGet,
			Path:        helpers.GET_AUTHOR_V1,
			Controller:  r.getOneAuthorHandler.GetOneAuthorAllBooks,
			Description: "Get One Author all Books",
			Type:        helpers.HANDLER_FUNC,
			Middlewares: []func(http.Handler) http.Handler{},
		},
	}
}
