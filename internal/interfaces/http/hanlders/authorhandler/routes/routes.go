package author_routes

import (
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/authorhandler"
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/authorhandler/middleware"
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/helpers"
	"github.com/andreis3/foodtosave-case/internal/util"
	"net/http"
)

type Routes struct {
	createAuthorHandler authorhandler.ICreateAuthorHandler
	getOneAuthorHandler authorhandler.IGetOneAuthorHandler
}

func NewAuthorRoutes(
	createAuthorHandler authorhandler.ICreateAuthorHandler,
	getOneAuthorHandler authorhandler.IGetOneAuthorHandler) *Routes {
	return &Routes{
		createAuthorHandler: createAuthorHandler,
		getOneAuthorHandler: getOneAuthorHandler,
	}
}

func (r *Routes) GroupRoutes() util.RouteType {
	return util.RouteType{
		{
			Method:      http.MethodPost,
			Path:        helpers.CREATE_AUTHOR_V1,
			Controller:  r.createAuthorHandler.CreateAuthorWithBooks,
			Description: "Create Author",
			Type:        util.HANDLER_FUNC,
			Middlewares: []func(http.Handler) http.Handler{
				group_middleware.ValidatePath,
			},
		},
		{
			Method:      http.MethodGet,
			Path:        helpers.GET_AUTHOR_V1,
			Controller:  r.getOneAuthorHandler.GetOneAuthorAllBooks,
			Description: "Get One Author",
			Type:        util.HANDLER_FUNC,
			Middlewares: []func(http.Handler) http.Handler{},
		},
	}
}
