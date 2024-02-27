package usecases

import (
	"context"
	"errors"
	"testing"

	"github.com/Lucasma95/golang-company-api/src/api/core/entities"
	contract "github.com/Lucasma95/golang-company-api/src/api/http/contracts/company"
	mocks "github.com/Lucasma95/golang-company-api/src/api/test/mocks/providers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CreateCompanySuccessfully(t *testing.T) {

	CompanyRepository := mocks.CompanyRepository{}
	CreateCompanyUsecase := CreateCompanyImpl{
		CompanyRepository: &CompanyRepository,
	}

	request := getCreateCompanyRequestMock()

	company := entities.NewCompany(request)

	CompanyRepository.On("Create", mock.Anything, &company).Return(nil)

	err := CreateCompanyUsecase.Execute(context.Background(), request)

	assert.Nil(t, err)
	CompanyRepository.AssertNumberOfCalls(t, "Create", 1)
}

func Test_CreateCompanyFails(t *testing.T) {

	CompanyRepository := mocks.CompanyRepository{}
	CreateCompanyUsecase := CreateCompanyImpl{
		CompanyRepository: &CompanyRepository,
	}

	request := getCreateCompanyRequestMock()

	company := entities.NewCompany(request)

	CompanyRepository.On("Create", mock.Anything, &company).Return(errors.New("a random error"))

	err := CreateCompanyUsecase.Execute(context.Background(), request)

	assert.Error(t, err)
	CompanyRepository.AssertNumberOfCalls(t, "Create", 1)
}

func getCreateCompanyRequestMock() *contract.CreateCompanyRequest {
	return &contract.CreateCompanyRequest{
		Name:        "name",
		Description: "description",
		CountryName: "country_name",
	}
}
