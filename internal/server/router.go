package server

import (
	"golang.frontdoorhome.com/personal-project/tiny-url/internal/config"
	"golang.frontdoorhome.com/personal-project/tiny-url/internal/handlers"
	"golang.frontdoorhome.com/personal-project/tiny-url/internal/literals"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
}

// NewRouter ...
func NewRouter() *Router {
	return &Router{mux.NewRouter()}
}

// InitializeRouter ...
func (r *Router) InitializeRouter(routerConfig *config.Config) {
	//	r.initializeMiddleware(routerConfig.WebServerConfig)
	r.initializeRoutes(routerConfig)
}

func (r *Router) initializeRoutes(routerConfig *config.Config) {
	s := (*r).PathPrefix(routerConfig.WebServerConfig.RoutePrefix).Subrouter()


	s.HandleFunc(literals.HealthcheckURI,
		handlers.HealthCheckHandler(routerConfig)).
		Methods(http.MethodGet).
		Name(literals.HealthcheckAPIName)
	s.HandleFunc(literals.TinyUrl,
		handlers.GetLongHandler(routerConfig)).
		Methods(http.MethodGet).
		Name(literals.GetLongUrlAPIName)
	s.HandleFunc(literals.TinyUrl,
		handlers.AddTinyUrlHandler(routerConfig)).
		Methods(http.MethodPost).
		Name(literals.CreateTinyUrlAPIName)

}

func (r *Router) initializeMiddleware(routerConfig *config.WebServerConfig) {

}
