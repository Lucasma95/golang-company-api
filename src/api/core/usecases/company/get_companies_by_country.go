package usecases

import (
	"context"

	"github.com/Lucasma95/golang-company-api/src/api/core/entities"
	"github.com/Lucasma95/golang-company-api/src/api/core/providers"
)

type GetCompaniesByCountry interface {
	Execute(ctx context.Context, countryName string) ([]entities.Company, error)
}

type GetCompaniesByCountryImpl struct {
	CompanyRepository providers.CompanyRepository
}

func (impl *GetCompaniesByCountryImpl) Execute(ctx context.Context, countryName string) ([]entities.Company, error) {

	companies, err := impl.CompanyRepository.GetByCountryName(ctx, countryName)
	if err != nil {
		return nil, err
	}

	return companies, nil
}
