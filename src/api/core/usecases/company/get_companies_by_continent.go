package usecases

import (
	"context"

	"github.com/Lucasma95/golang-company-api/src/api/core/entities"
	"github.com/Lucasma95/golang-company-api/src/api/core/providers"
)

type GetCompaniesByContinent interface {
	Execute(ctx context.Context, continent string) ([]entities.Company, error)
}

type GetCompaniesByContinentImpl struct {
	CompanyRepository providers.CompanyRepository
}

func (impl *GetCompaniesByContinentImpl) Execute(ctx context.Context, continent string) ([]entities.Company, error) {

	companies, err := impl.CompanyRepository.GetByContinent(ctx, continent)
	if err != nil {
		return nil, err
	}

	return companies, nil
}
