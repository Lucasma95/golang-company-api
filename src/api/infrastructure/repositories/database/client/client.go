package database

import (
	"fmt"
	"log"

	"github.com/Lucasma95/golang-company-api/src/api/application/config"
	"github.com/Lucasma95/golang-company-api/src/api/core/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/tracing"
)

func New(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		config.DB.Host,
		config.DB.Port,
		config.DB.User,
		config.DB.Name,
		config.DB.Password,
	)

	client, err := gorm.Open(postgres.Open(dsn),
		&gorm.Config{Logger: gormLogger.Default.LogMode(gormLogger.Silent)})
	if err != nil {
		panic(err)
	}

	if config.DB.AutomigrationEnabled {
		err = client.AutoMigrate(&entities.Company{}, &entities.Country{})
		if err != nil {
			log.Println("error automigrating the entities")
		}
	}

	if err := client.Use(tracing.NewPlugin(tracing.WithDBName(fmt.Sprintf("%s-db", config.DB.Name)))); err != nil {
		log.Println("error adding tracer to gorm")
	}

	return client
}
