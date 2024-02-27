package entities

import (
	contract "github.com/Lucasma95/golang-company-api/src/api/http/contracts/country"
)

type Country struct {
	Name      string `gorm:"type:varchar;primaryKey"`
	Continent string
}

func NewCountry(request *contract.CreateCountryRequest) Country {
	return Country{
		Name:      request.Name,
		Continent: request.Continent,
	}
}
