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

func Test_GetCompanyByContinentSuccessfully(t *testing.T) {

	CompanyRepository := mocks.CompanyRepository{}
	usecase := GetCompaniesByContinentImpl{
		CompanyRepository: &CompanyRepository,
	}

	continent := "America"

	companies := []entities.Company{getMockCompany("id")}

	CompanyRepository.On("GetByContinent", mock.Anything, continent).Return(companies, nil)

	result, err := usecase.Execute(context.Background(), continent)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
	CompanyRepository.AssertNumberOfCalls(t, "GetByContinent", 1)
}

func Test_GetCompanyByContinentFails(t *testing.T) {

	CompanyRepository := mocks.CompanyRepository{}
	usecase := GetCompaniesByContinentImpl{
		CompanyRepository: &CompanyRepository,
	}

	continent := "America"

	CompanyRepository.On("GetByContinent", mock.Anything, continent).Return(nil, errors.New("a random error"))

	result, err := usecase.Execute(context.Background(), continent)

	assert.Error(t, err)
	assert.Equal(t, 0, len(result))
	CompanyRepository.AssertNumberOfCalls(t, "GetByContinent", 1)
}
