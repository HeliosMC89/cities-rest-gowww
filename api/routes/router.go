package routes

import (
	"errors"
	"log"
	"net/http"

	"github.com/gowww/router"
	"github.com/heliosmc89/api-rest-gowww/api/handlers"
	"github.com/heliosmc89/api-rest-gowww/api/middlewares"
	"github.com/heliosmc89/api-rest-gowww/helpers"
	"github.com/jmoiron/sqlx"
)

func NewRouter(logging *log.Logger, db *sqlx.DB) *router.Router {
	cityHandler := handlers.NewCityHandler(logging, db)
	routes := Routes{
		Route{"CityAll", "GET", "/cities", cityHandler.GetAllCities},
		Route{"CityAdd", "POST", "/cities", cityHandler.CreateCity},
		Route{"CityByID", "GET", `/cities/:id:^\d+$`, cityHandler.GetCityByID},
		Route{"CityUpdate", "PUT", `/cities/:id:^\d+$`, cityHandler.UpdateCity},
		Route{"CityDelete", "DELETE", `/cities/:id:^\d+$`, cityHandler.DeleteCity},
	}
	rt := router.New()
	rt.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := errors.New("Resource not found")
		helpers.WriteJSON(w, http.StatusNotFound, helpers.JSONResponse{
			Error: true,
			Msg:   err.Error(),
			Code:  http.StatusNotFound,
			Data:  nil,
		})
	})
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = middlewares.Logger(handler, route.Name, logging)
		rt.Handle(route.Method, route.Path, handler)
	}

	return rt
}
