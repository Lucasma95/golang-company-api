package database

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Lucasma95/golang-company-api/src/api/core/entities"
	mocks "github.com/Lucasma95/golang-company-api/src/api/test/mocks/database"
	"github.com/stretchr/testify/assert"
)

const (
	selectCompany = `SELECT \* FROM "companies".*`
	selectCountry = `SELECT \* FROM "countries".*`
	insertCompany = `INSERT INTO "companies".*`
)

func Test_CreateCompanySuccessfully(t *testing.T) {
	mockDB, DB := mocks.New()
	mockDB.MatchExpectationsInOrder(false)

	company := getCompanyMock()

	repository := CompanyRepository{
		Client: DB.Debug(),
	}

	mockDB.ExpectBegin()
	mockDB.ExpectExec(insertCompany).WillReturnResult(sqlmock.NewResult(0, 1))
	mockDB.ExpectCommit()

	err := repository.Create(context.Background(), company)

	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func Test_GetCompanyByID(t *testing.T) {
	mockDB, DB := mocks.New()
	mockDB.MatchExpectationsInOrder(false)

	repository := CompanyRepository{
		Client: DB.Debug(),
	}

	companyID := "company-id"

	subMockRows := sqlmock.NewRows([]string{"id", "description", "country_name"}).
		AddRow(companyID, "description", "Argentina")

	mockDB.ExpectQuery(selectCompany).WithArgs(companyID).WillReturnRows(subMockRows)

	mockCountryRows := sqlmock.NewRows([]string{"name", "continent"}).
		AddRow("Argentina", "America")

	mockDB.ExpectQuery(selectCountry).
		WithArgs("Argentina").
		WillReturnRows(mockCountryRows)

	company, err := repository.GetByID(context.Background(), companyID)

	assert.NoError(t, err)
	assert.NotNil(t, company)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func Test_GetCompaniesByCountry(t *testing.T) {
	mockDB, DB := mocks.New()
	mockDB.MatchExpectationsInOrder(false)

	repository := CompanyRepository{
		Client: DB.Debug(),
	}

	country := "Argentina"

	subMockRows := sqlmock.NewRows([]string{"id", "description", "country_name"}).
		AddRow("company-id", "description", "Argentina")

	mockDB.ExpectQuery(selectCompany).WithArgs(country).WillReturnRows(subMockRows)

	mockCountryRows := sqlmock.NewRows([]string{"name", "continent"}).
		AddRow(country, "America")

	mockDB.ExpectQuery(selectCountry).
		WithArgs(country).
		WillReturnRows(mockCountryRows)

	companies, err := repository.GetByCountryName(context.Background(), country)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(companies))
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func Test_GetCompaniesByContinent(t *testing.T) {
	mockDB, DB := mocks.New()
	mockDB.MatchExpectationsInOrder(false)

	repository := CompanyRepository{
		Client: DB.Debug(),
	}

	continent := "America"

	subMockRows := sqlmock.NewRows([]string{"id", "description", "country_name"}).
		AddRow("company-id", "description", "Argentina")

	mockDB.ExpectQuery(selectCompany).WithArgs(continent).WillReturnRows(subMockRows)

	mockCountryRows := sqlmock.NewRows([]string{"name", "continent"}).
		AddRow("Argentina", continent)

	mockDB.ExpectQuery(selectCountry).
		WithArgs("Argentina").
		WillReturnRows(mockCountryRows)

	companies, err := repository.GetByContinent(context.Background(), continent)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(companies))
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func getCompanyMock() *entities.Company {
	return &entities.Company{
		ID:          "id",
		Name:        "company_name",
		Description: "description",
		CountryName: "Argentina",
		Country: entities.Country{
			Name:      "Argentina",
			Continent: "America",
		},
	}
}
