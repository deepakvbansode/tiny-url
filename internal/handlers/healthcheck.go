package handlers

import (
	"fmt"
	"golang.frontdoorhome.com/personal-project/tiny-url/internal/config"
	"golang.frontdoorhome.com/personal-project/tiny-url/internal/services"
	"golang.frontdoorhome.com/personal-project/tiny-url/internal/utils"
	"net/http"
)

func HealthCheckHandler(config *config.Config) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("heathcheck called")
		healthService := services.GetServiceFactory().GetHealthService()

		healthService.PrintHealth()
		utils.WriteResponse(writer, http.StatusOK, config)
	}
}
