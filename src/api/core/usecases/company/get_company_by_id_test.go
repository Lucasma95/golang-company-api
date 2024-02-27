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

func Test_GetCompanyByIDSuccessfully(t *testing.T) {

	CompanyRepository := mocks.CompanyRepository{}
	usecase := GetCompanyByIDImpl{
		CompanyRepository: &CompanyRepository,
	}

	companyID := "id"

	company := getMockCompany(companyID)

	CompanyRepository.On("GetByID", mock.Anything, companyID).Return(&company, nil)

	result, err := usecase.Execute(context.Background(), companyID)

	assert.Nil(t, err)
	assert.Equal(t, companyID, result.ID)
	CompanyRepository.AssertNumberOfCalls(t, "GetByID", 1)
}

func Test_GetCompanyByIDFails(t *testing.T) {

	CompanyRepository := mocks.CompanyRepository{}
	usecase := GetCompanyByIDImpl{
		CompanyRepository: &CompanyRepository,
	}

	companyID := "id"

	CompanyRepository.On("GetByID", mock.Anything, companyID).Return(nil, errors.New("a random error"))

	result, err := usecase.Execute(context.Background(), companyID)

	assert.Error(t, err)
	assert.Nil(t, result)
	CompanyRepository.AssertNumberOfCalls(t, "GetByID", 1)
}

func getMockCompany(id string) entities.Company {
	return entities.Company{
		ID: id,
		Name: "name",
		Description: "description",
		CountryName: "Argentina",
		Country: entities.Country{
			Name: "Argentina",
			Continent: "America",
		},
	}
}