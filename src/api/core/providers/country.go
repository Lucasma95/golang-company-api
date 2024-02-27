package providers

import (
	"context"

	"github.com/Lucasma95/golang-company-api/src/api/core/entities"
)

type CountryRepository interface {
	Create(ctx context.Context, country *entities.Country) error
}
