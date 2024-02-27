package api

import (
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
)

func init() {
	_ = os.Setenv("GIN_MODE", "test")
}

func Test_GetCompanyByIdSuccessfully(t *testing.T) {

	companyID := "company_id_123"

	expectedCompany := entities.Company{ID: companyID}

	getCompanyUsecase := mocks.GetCompanyByID{}
	getCompanyByIDHandler := GetCompanyByID{GetCompanyByIDUsecase: &getCompanyUsecase}

	getCompanyUsecase.On("Execute", mock.Anything, companyID).Return(&expectedCompany, nil)

	r := gin.Default()
	groupAPI := r.Group("/api/v1")
	groupAPI.GET("company/:company_id", getCompanyByIDHandler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/company/company_id_123", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	getCompanyUsecase.AssertNumberOfCalls(t, "Execute", 1)
}

func Test_GetCompanyByIdFails(t *testing.T) {

	companyID := "company_id_123"

	getCompanyUsecase := mocks.GetCompanyByID{}
	getCompanyByIDHandler := GetCompanyByID{GetCompanyByIDUsecase: &getCompanyUsecase}

	getCompanyUsecase.On("Execute", mock.Anything, companyID).Return(nil, errors.New("a random error"))

	r := gin.Default()
	groupAPI := r.Group("/api/v1")
	groupAPI.GET("company/:company_id", getCompanyByIDHandler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/company/company_id_123", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	getCompanyUsecase.AssertNumberOfCalls(t, "Execute", 1)
}

func Test_GetCompanyByIdReturnsNotFound(t *testing.T) {

	companyID := "company_id_123"

	getCompanyUsecase := mocks.GetCompanyByID{}
	getCompanyByIDHandler := GetCompanyByID{GetCompanyByIDUsecase: &getCompanyUsecase}

	getCompanyUsecase.On("Execute", mock.Anything, companyID).Return(nil, nil)

	r := gin.Default()
	groupAPI := r.Group("/api/v1")
	groupAPI.GET("company/:company_id", getCompanyByIDHandler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/company/company_id_123", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	getCompanyUsecase.AssertNumberOfCalls(t, "Execute", 1)
}
