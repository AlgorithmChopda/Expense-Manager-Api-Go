package routes

import (
	"github.com/gorilla/mux"
)

func UseRoute(prefix string, subRouterFunc func(router *mux.Router), router *mux.Router) {
	subRouter := router.PathPrefix(prefix).Subrouter()
	subRouterFunc(subRouter)
}

func Route() *mux.Router {
	router := mux.NewRouter()
	UseRoute("/user", UserRoutes, router)
	UseRoute("/test", TestRoutes, router)
	return router
}
