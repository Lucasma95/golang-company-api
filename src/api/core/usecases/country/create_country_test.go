package usecases

import (
	"context"
	"errors"
	"testing"

	"github.com/Lucasma95/golang-company-api/src/api/core/entities"
	contract "github.com/Lucasma95/golang-company-api/src/api/http/contracts/country"
	mocks "github.com/Lucasma95/golang-company-api/src/api/test/mocks/providers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CreateCountrySuccessfully(t *testing.T) {

	countryRepository := mocks.CountryRepository{}
	CreateCompanyUsecase := CreateCountryImpl{
		CountryRepository: &countryRepository,
	}

	request := getCreateCountryRequestMock()

	country := entities.NewCountry(request)

	countryRepository.On("Create", mock.Anything, &country).Return(nil)

	err := CreateCompanyUsecase.Execute(context.Background(), request)

	assert.Nil(t, err)
	countryRepository.AssertNumberOfCalls(t, "Create", 1)
}

func Test_CreateCountryFails(t *testing.T) {

	countryRepository := mocks.CountryRepository{}
	CreateCompanyUsecase := CreateCountryImpl{
		CountryRepository: &countryRepository,
	}

	request := getCreateCountryRequestMock()

	country := entities.NewCountry(request)

	countryRepository.On("Create", mock.Anything, &country).Return(errors.New("a random error"))

	err := CreateCompanyUsecase.Execute(context.Background(), request)

	assert.Error(t, err)
	countryRepository.AssertNumberOfCalls(t, "Create", 1)
}

func getCreateCountryRequestMock() *contract.CreateCountryRequest {
	return &contract.CreateCountryRequest{
		Name:      "name",
		Continent: "continent",
	}
}
