package database

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func New() (sqlmock.Sqlmock, *gorm.DB) {
	db, mockDB, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	if err != nil {
		log.Fatalf("can't create sqlmock: %s", err)
	}

	clientDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db, PreferSimpleProtocol: true}),
		&gorm.Config{Logger: gormLogger.Default.LogMode(gormLogger.Info)})

	if err != nil {
		log.Fatalf("can't open gorm connection: %s", err)
	}

	clientDB.Set("gorm:update_column", true)

	return mockDB, clientDB
}
