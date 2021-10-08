package services

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.frontdoorhome.com/personal-project/tiny-url/internal/models"
	"golang.frontdoorhome.com/personal-project/tiny-url/internal/repositories"
)

type UrlService interface {
	GetUrl(tinyUrl string) string
	AddTinyUrl(url *models.Url) bool
}

type url struct {
	dao repositories.DaoFactory
}

func (e *url) GetUrl(tinyUrl string) string {
	fmt.Println("Inside URL service", e.dao)
	//redis first and then db
	longUrl := e.dao.GetUrlDao().GetLongUrl(tinyUrl)
	return longUrl
}

func (e *url)  AddTinyUrl(url *models.Url) bool {

	log.Println("Inside url service", url)
	//GetServiceFactory().GetHealthService().PrintHealth()
//	_ := e.dao.GetEmployeeDao().GetEmployee()
	return true
}


func getUrlService(dao repositories.DaoFactory) UrlService {
	return &url{
		dao: dao,
	}
}
