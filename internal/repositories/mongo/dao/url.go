package dao

import (
	"fmt"
	"golang.frontdoorhome.com/personal-project/tiny-url/internal/models"
	"golang.frontdoorhome.com/personal-project/tiny-url/internal/repositories"
)

/**
it will be like service factoryer
it will also connect the server
basically init dao will
	connect the server and build the singleton instance of dao Factoryer
*/

type urlDao struct {
	mongo interface{}
}

func (e *urlDao) GetLongUrl(tinyUrl string) string {
	fmt.Println("Getting data from mongo:", e.mongo)
	return "http://www.xyz.com/very-long-url"
}


func (e *urlDao) InsertTinyUrl(url models.Url) bool {
	fmt.Println("Getting data from mongo:", e.mongo)
	return true
}

func getUrlDao(mongo interface{}) repositories.UrlDao {
	return &urlDao{
		mongo: mongo,
	}
}
