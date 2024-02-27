package providers

import (
	"context"

	"github.com/Lucasma95/golang-company-api/src/api/core/entities"
)

type CompanyRepository interface {
	GetByID(ctx context.Context, id string) (*entities.Company, error)
	GetByCountryName(ctx context.Context, countryName string) ([]entities.Company, error)
	GetByContinent(ctx context.Context, continent string) ([]entities.Company, error)
	Create(ctx context.Context, company *entities.Company) error
}
