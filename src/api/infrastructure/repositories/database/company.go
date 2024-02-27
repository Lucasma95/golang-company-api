package database

import (
	"context"

	"github.com/Lucasma95/golang-company-api/src/api/core/entities"
	"gorm.io/gorm"
)

type CompanyRepository struct {
	Client *gorm.DB
}

func (repo *CompanyRepository) GetByContinent(ctx context.Context, continent string) ([]entities.Company, error) {
    var companies []entities.Company
    result := repo.Client.Preload("Country").Where("continent = ?", continent).Find(&companies)
    if result.Error != nil {
        return nil, result.Error
    }
    return companies, nil
}

func (repo *CompanyRepository) GetByCountryName(ctx context.Context, countryName string) ([]entities.Company, error) {
	return getEntitiesByParams[entities.Company](ctx, repo.Client, map[string]any{"country_name": countryName}, "Country")
}

func (repo *CompanyRepository) GetByID(ctx context.Context, id string) (*entities.Company, error) {
	return getByID[entities.Company](ctx, repo.Client, id, "Country")
}

func (repo *CompanyRepository) Create(ctx context.Context, company *entities.Company) error {
	return create(ctx, repo.Client, company, "Country")
}
