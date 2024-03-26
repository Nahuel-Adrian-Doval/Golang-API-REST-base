package pkg

import (
	"fmt"

	"github.com/Nahuel-Adrian-Doval/Golang-API-REST-base/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DataBaseInstance *gorm.DB

func DBConnection(conf models.DBConfig) (db *gorm.DB, err error) {

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", conf.DB_HOST, conf.DB_USER, conf.DB_PASSWORD, conf.DB_NAME, conf.DB_PORT, conf.DB_SSLMODE, conf.DB_TIMEZONE)

	if DataBaseInstance, err = gorm.Open(postgres.Open(dns), &gorm.Config{}); err != nil {
		return db, err
	}

	return db, nil
}
