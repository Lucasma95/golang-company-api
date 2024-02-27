package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Lucasma95/golang-company-api/src/api/core/entities"
	mocks "github.com/Lucasma95/golang-company-api/src/api/test/mocks/usecases"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func init() {
	_ = os.Setenv("GIN_MODE", "test")
}

func Test_GetCompaniesByContinentSuccessfully(t *testing.T) {

	continent := "America"

	expectedCompanies := []entities.Company{{Country: entities.Country{Continent: continent}}}

	getCompanyUsecase := mocks.GetCompanyByContinent{}
	getCompaniesByContinentHandler := GetCompaniesByContinent{GetCompaniesByContinentUsecase: &getCompanyUsecase}

	getCompanyUsecase.On("Execute", mock.Anything, continent).Return(expectedCompanies, nil)

	r := gin.Default()
	groupAPI := r.Group("/api/v1")
	groupAPI.GET("company/continent/:continent", getCompaniesByContinentHandler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/company/continent/America", nil)
	r.ServeHTTP(w, req)
	
	var companiesRetrieved []entities.CompanyDTO
	
	decoder := json.NewDecoder(w.Result().Body)
	
	err := decoder.Decode(&companiesRetrieved)
	require.Nil(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, len(companiesRetrieved), 1)
	getCompanyUsecase.AssertNumberOfCalls(t, "Execute", 1)
}

func Test_GetCompanyByContinentFails(t *testing.T) {

	continent := "America"

	getCompanyUsecase := mocks.GetCompanyByContinent{}
	getCompaniesByContinentHandler := GetCompaniesByContinent{GetCompaniesByContinentUsecase: &getCompanyUsecase}

	getCompanyUsecase.On("Execute", mock.Anything, continent).Return(nil, errors.New("a random error"))

	r := gin.Default()
	groupAPI := r.Group("/api/v1")
	groupAPI.GET("company/continent/:continent", getCompaniesByContinentHandler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/company/continent/America", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	getCompanyUsecase.AssertNumberOfCalls(t, "Execute", 1)
}

func Test_GetCompanyByContinentReturnsNoCompanies(t *testing.T) {

	continent := "America"

	expectedCompanies := []entities.Company{}

	getCompanyUsecase := mocks.GetCompanyByContinent{}
	getCompaniesByContinentHandler := GetCompaniesByContinent{GetCompaniesByContinentUsecase: &getCompanyUsecase}

	getCompanyUsecase.On("Execute", mock.Anything, continent).Return(expectedCompanies, nil)

	r := gin.Default()
	groupAPI := r.Group("/api/v1")
	groupAPI.GET("company/continent/:continent", getCompaniesByContinentHandler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/company/continent/America", nil)
	r.ServeHTTP(w, req)
	
	var companiesRetrieved []entities.CompanyDTO
	
	decoder := json.NewDecoder(w.Result().Body)
	
	err := decoder.Decode(&companiesRetrieved)
	require.Nil(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, len(companiesRetrieved), 0)
	getCompanyUsecase.AssertNumberOfCalls(t, "Execute", 1)
}
