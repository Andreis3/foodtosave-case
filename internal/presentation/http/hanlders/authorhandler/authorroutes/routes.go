package authorroutes

import (
	"net/http"

	"github.com/andreis3/foodtosave-case/internal/infra/common/logger"
	"github.com/andreis3/foodtosave-case/internal/presentation/http/hanlders/authorhandler"
	"github.com/andreis3/foodtosave-case/internal/presentation/http/helpers"
	"github.com/andreis3/foodtosave-case/internal/presentation/http/middleware"
)

type Routes struct {
	createAuthorHandler authorhandler.CreateAuthorWithBooksHandler
	getOneAuthorHandler authorhandler.GetOneAuthorAllBooksHandler
	one                 *middleware.MiddlewareOne
	two                 *middleware.MiddlewareTwo
}

func NewAuthorRoutes(
	createAuthorHandler authorhandler.CreateAuthorWithBooksHandler,
	getOneAuthorHandler authorhandler.GetOneAuthorAllBooksHandler) *Routes {
	log := logger.NewLogger()
	return &Routes{
		createAuthorHandler: createAuthorHandler,
		getOneAuthorHandler: getOneAuthorHandler,
		one:                 middleware.NewMiddlewareOne(log),
		two:                 middleware.NewMiddlewareTwo(log),
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
				r.two.Two,
				r.one.One,
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
