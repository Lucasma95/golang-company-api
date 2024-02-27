package database

import (
	"context"

	"github.com/Lucasma95/golang-company-api/src/api/core/entities"
	"gorm.io/gorm"
)

type CountryRepository struct {
	Client *gorm.DB
}

func (repo *CountryRepository) Create(ctx context.Context, country *entities.Country) error {
	return create(ctx, repo.Client, country)
}
