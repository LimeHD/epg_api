package service

import (
	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/oschwald/geoip2-golang"
	"sync"
)

var once sync.Once
var instance *Service

type Service struct {
	Database  *dbx.DB
	GeoReader *geoip2.Reader
	// .. config, cache & etc
}

func (service *Service) ConnectDatabase(dataSource string) {
	var err error
	service.Database, err = dbx.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}
}

func GetInstance() *Service {
	once.Do(func() {
		instance = &Service{}
	})

	return instance
}
