package dao

import (
	"golang.frontdoorhome.com/personal-project/tiny-url/internal/repositories"
	"sync"
)

type daoFactory struct {
	mongo interface{}
}

var daoFactoryInstance repositories.DaoFactory
var once sync.Once

func InitDao(mongo interface{}) repositories.DaoFactory {
	once.Do(func() {
		daoFactoryInstance = &daoFactory{
			mongo: mongo,
		}
	})
	return daoFactoryInstance
}

func GetDaoFactory() repositories.DaoFactory {
	return daoFactoryInstance
}

func (df *daoFactory) GetUrlDao() repositories.UrlDao {
	return getUrlDao(df.mongo)
}

//
//func ConnectMongoDB(conf *mongo.Config) (mongo.DatabaseHelper, error) {
//	clOpts := []*options.ClientOptions{}
//	dbOpts := []*options.DatabaseOptions{}
//
//	// connect to database
//	db, err := mongo.NewDatabase(conf, clOpts, dbOpts)
//	if err != nil {
//		log.Errorf("MongoDB initialization failed, reason: %v", err.Error())
//		return nil, err
//	}
//
//	// set db handle in router config struct
//	return db, nil
//}
