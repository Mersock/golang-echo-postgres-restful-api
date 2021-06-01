package storage

import (
	"log"

	"github.com/Mersock/golang-echo-postgres-restful-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB(params ...string) *gorm.DB {
	var err error
	connString := config.GetPostgresConn()

	log.Print(connString)

	DB, err = gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
