package server

import (
	"fmt"
	muxHandler "github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
	"golang.frontdoorhome.com/personal-project/tiny-url/internal/config"
	"golang.frontdoorhome.com/personal-project/tiny-url/internal/repositories"
	"golang.frontdoorhome.com/personal-project/tiny-url/internal/repositories/mongo/dao"
	"golang.frontdoorhome.com/personal-project/tiny-url/internal/services"
	"net/http"
)

//Server ...
type Server struct {
	Configuration *config.WebServerConfig
	Router        *Router
}

//NewServer ...
func NewServer(config *config.WebServerConfig) *Server {
	server := &Server{
		Configuration: config,
		Router:        NewRouter(),
	}

	return server
}

func RunServer() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.Println("Running server")
	webServerConfig, err := config.FromEnv()
	if err != nil {
		return err
	}
	fmt.Println("config", webServerConfig)
	//Todo: start mongo

	server := NewServer(webServerConfig)

	config := config.Config{
		WebServerConfig: webServerConfig,
	}
	var daoFactory repositories.DaoFactory
	daoFactory = dao.InitDao("pass the cred here")

	services.InitServiceFactory(daoFactory)

	server.Router.InitializeRouter(&config)
	cors := muxHandler.CORS(muxHandler.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), muxHandler.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), muxHandler.AllowedOrigins([]string{"*"}))
	log.Println("starting server")
	err = http.ListenAndServe(":"+webServerConfig.Port, cors(server.Router))
	if err != nil {
		return err
	}

	return nil
}
