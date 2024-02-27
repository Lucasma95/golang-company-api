package usecases

import (
	"context"

	"github.com/Lucasma95/golang-company-api/src/api/core/entities"
	"github.com/Lucasma95/golang-company-api/src/api/core/providers"
	contract "github.com/Lucasma95/golang-company-api/src/api/http/contracts/country"
)

type CreateCountry interface {
	Execute(ctx context.Context, request *contract.CreateCountryRequest) error
}

type CreateCountryImpl struct {
	CountryRepository providers.CountryRepository
}

func (impl *CreateCountryImpl) Execute(ctx context.Context, request *contract.CreateCountryRequest) error {

	country := entities.NewCountry(request)

	if err := impl.CountryRepository.Create(ctx, &country); err != nil {
		return err
	}

	return nil
}
