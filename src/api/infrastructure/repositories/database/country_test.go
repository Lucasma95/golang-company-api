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
	insertCountry = `INSERT INTO "countries".*`
)

func Test_CreateCountrySuccessfully(t *testing.T) {
	mockDB, DB := mocks.New()
	mockDB.MatchExpectationsInOrder(false)

	country := getCountryMock()

	repository := CountryRepository{
		Client: DB.Debug(),
	}

	mockDB.ExpectBegin()
	mockDB.ExpectExec(insertCountry).WillReturnResult(sqlmock.NewResult(0, 1))
	mockDB.ExpectCommit()

	err := repository.Create(context.Background(), country)

	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func getCountryMock() *entities.Country {
	return &entities.Country {
		Name:      "Argentina",
		Continent: "America",
	}
}
