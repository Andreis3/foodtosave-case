package author_routes

import (
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/author"
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/hanlders/author/middleware"
	"github.com/andreis3/foodtosave-case/internal/interfaces/http/helpers"
	"github.com/andreis3/foodtosave-case/internal/util"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

type Routes struct {
	createAuthorHandler handler_author.ICreateAuthorHandler
	getOneAuthorHandler handler_author.IGetOneAuthorHandler
}

func NewAuthorRoutes(
	createAuthorHandler handler_author.ICreateAuthorHandler,
	getOneAuthorHandler handler_author.IGetOneAuthorHandler) *Routes {
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
			Controller:  r.createAuthorHandler.CreateAuthor,
			Description: "Create Author",
			Type:        util.HANDLER_FUNC,
			Middlewares: []func(http.Handler) http.Handler{
				group_middleware.ValidatePath,
				middleware.Logger,
			},
		},
		{
			Method:      http.MethodGet,
			Path:        helpers.GET_AUTHOR_V1,
			Controller:  r.getOneAuthorHandler.GetOneAuthor,
			Description: "Get One Author",
			Type:        util.HANDLER_FUNC,
			Middlewares: []func(http.Handler) http.Handler{},
		},
	}
}
