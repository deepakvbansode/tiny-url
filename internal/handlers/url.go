package handlers

import (
	"encoding/json"
	"fmt"
	"golang.frontdoorhome.com/personal-project/tiny-url/internal/config"
	"golang.frontdoorhome.com/personal-project/tiny-url/internal/models"
	"golang.frontdoorhome.com/personal-project/tiny-url/internal/services"
	"golang.frontdoorhome.com/personal-project/tiny-url/internal/utils"
	"io/ioutil"
	"net/http"
)

func GetLongHandler(config *config.Config) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("employee called")
		employeeService := services.GetServiceFactory().GetUrlService()

		tinyURL := "need to get it from request"
		employee := employeeService.GetUrl(tinyURL)
		utils.WriteResponse(writer, http.StatusOK, employee)
	}
}

func AddTinyUrlHandler(config *config.Config) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		url := models.Url{}

		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			utils.WriteResponse(writer, http.StatusBadRequest, err.Error())
			return
		}

		err = json.Unmarshal(body, &url)
		if err != nil {
			utils.WriteResponse(writer, http.StatusBadRequest, err.Error())
			return
		}

		res, err := url, nil
		if err != nil {
			utils.WriteResponse(writer, http.StatusBadRequest, err.Error())
			return
		}
		utils.WriteResponse(writer, http.StatusOK, res)
	}
}
