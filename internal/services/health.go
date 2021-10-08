package services

import (
	"fmt"
	"golang.frontdoorhome.com/personal-project/tiny-url/internal/repositories"
	"sync"

)

/*
	this service is single ton service example
*/

type HealthService interface {
	PrintHealth() error
}

type health struct {
	dao repositories.DaoFactory
}

var instanceInstance HealthService
var healthOnce sync.Once

func getHealthService(dao repositories.DaoFactory) HealthService {
	healthOnce.Do(func() {
		instanceInstance = &health{dao: dao}
	})
	return instanceInstance
}

func (h *health) PrintHealth() error {
	fmt.Println("I am healthy from services...")
	fmt.Println("I will use this dao:", h.dao)

	return nil
}
