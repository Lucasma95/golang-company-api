package usecases

import (
	"context"

	"github.com/Lucasma95/golang-company-api/src/api/core/entities"
	"github.com/Lucasma95/golang-company-api/src/api/core/providers"
)

type GetCompanyByID interface {
	Execute(ctx context.Context, id string) (*entities.Company, error)
}

type GetCompanyByIDImpl struct {
	CompanyRepository providers.CompanyRepository
}

func (impl *GetCompanyByIDImpl) Execute(ctx context.Context, id string) (*entities.Company, error) {

	company, err := impl.CompanyRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return company, nil
}
