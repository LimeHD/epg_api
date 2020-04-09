package service

import (
	"github.com/bugsnag/bugsnag-go"
	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/savsgio/go-logger"
	"sync"
)

var once sync.Once
var instance *Service

type Service struct {
	Database        *dbx.DB
	BugsnagNotifier BugsnagService
	// .. config, cache & etc
}

type BugsnagService struct {
	NotifierInstance *bugsnag.Notifier
}

func (bugsnag *BugsnagService) Notify(err error, rawData ...interface{}) {
	err = bugsnag.NotifierInstance.Notify(err, rawData)

	if err != nil {
		logger.Warningf("Bugsnag warning: cant't send event message")
	}
}

func (service *Service) ConnectDatabase(dataSource string) {
	var err error
	service.Database, err = dbx.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}
}

func (service *Service) RegisterBugsnagNotifier(apikey string) {
	service.BugsnagNotifier.NotifierInstance = bugsnag.New(bugsnag.Configuration{
		APIKey:          apikey,
		AppVersion:      "0.0.1",
		ProjectPackages: []string{"main", "https://github.com/LimeHD/epg_api"},
	})
}

func GetInstance() *Service {
	once.Do(func() {
		instance = &Service{}
	})

	return instance
}
