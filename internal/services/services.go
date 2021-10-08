package services

import (
	"golang.frontdoorhome.com/personal-project/tiny-url/internal/repositories"
	"sync"
)

type ServiceFactory interface {
	GetHealthService() HealthService
	GetUrlService() UrlService
}

type serviceFactory struct {
	dao repositories.DaoFactory
}

var serviceFactoryInstance ServiceFactory
var once sync.Once

func InitServiceFactory(dao repositories.DaoFactory) ServiceFactory {
	once.Do(func() {
		serviceFactoryInstance = &serviceFactory{
			dao: dao,
		}
	})
	return serviceFactoryInstance
}

func GetServiceFactory() ServiceFactory {

	return serviceFactoryInstance
}

func (sf *serviceFactory) GetHealthService() HealthService {
	return getHealthService(sf.dao)
}

func (sf *serviceFactory) GetUrlService() UrlService {
	return getUrlService(sf.dao)
}
