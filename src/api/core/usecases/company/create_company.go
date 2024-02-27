package usecases

import (
	"context"

	"github.com/Lucasma95/golang-company-api/src/api/core/entities"
	"github.com/Lucasma95/golang-company-api/src/api/core/providers"
	contract "github.com/Lucasma95/golang-company-api/src/api/http/contracts/company"
)

type CreateCompany interface {
	Execute(ctx context.Context, request *contract.CreateCompanyRequest) error
}

type CreateCompanyImpl struct {
	CompanyRepository providers.CompanyRepository
}

func (impl *CreateCompanyImpl) Execute(ctx context.Context, request *contract.CreateCompanyRequest) error {

	company := entities.NewCompany(request)

	if err := impl.CompanyRepository.Create(ctx, &company); err != nil {
		return err
	}

	return nil
}
