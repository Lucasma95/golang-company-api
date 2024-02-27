package usecases

import (
	"context"
	"errors"
	"testing"

	"github.com/Lucasma95/golang-company-api/src/api/core/entities"
	mocks "github.com/Lucasma95/golang-company-api/src/api/test/mocks/providers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_GetCompanyByCountrySuccessfully(t *testing.T) {

	CompanyRepository := mocks.CompanyRepository{}
	usecase := GetCompaniesByCountryImpl{
		CompanyRepository: &CompanyRepository,
	}

	country := "Argentina"

	companies := []entities.Company{getMockCompany("id")}

	CompanyRepository.On("GetByCountryName", mock.Anything, country).Return(companies, nil)

	result, err := usecase.Execute(context.Background(), country)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
	CompanyRepository.AssertNumberOfCalls(t, "GetByCountryName", 1)
}

func Test_GetCompanyByCountryFails(t *testing.T) {

	CompanyRepository := mocks.CompanyRepository{}
	usecase := GetCompaniesByCountryImpl{
		CompanyRepository: &CompanyRepository,
	}

	country := "Argentina"

	CompanyRepository.On("GetByCountryName", mock.Anything, country).Return(nil, errors.New("a random error"))

	result, err := usecase.Execute(context.Background(), country)

	assert.Error(t, err)
	assert.Equal(t, 0, len(result))
	CompanyRepository.AssertNumberOfCalls(t, "GetByCountryName", 1)
}