package repositories

import "golang.frontdoorhome.com/personal-project/tiny-url/internal/models"

type UrlDao interface {
	GetLongUrl(tinyUrl string) string
	InsertTinyUrl(url models.Url) bool
}

type DaoFactory interface {
	GetUrlDao() UrlDao
}
