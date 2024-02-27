package api

import (
	"bytes"
	"errors"
	"os"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	contracts "github.com/Lucasma95/golang-company-api/src/api/http/contracts/company"
	mocks "github.com/Lucasma95/golang-company-api/src/api/test/mocks/usecases"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func init() {
	_ = os.Setenv("GIN_MODE", "test")
}

func Test_CreateCompanySuccessfully(t *testing.T) {

	createCompanyUseCase := mocks.CreateCompany{}
	createCompanyHandler := CreateCompany{CreateCompanyUsecase: &createCompanyUseCase}

	request := getCreateCompanyRequestMock()
	jsonRequest := MarshallRequest(request)

	createCompanyUseCase.On("Execute", mock.Anything, &request).Return(nil)
	r := gin.Default()
	groupAPI := r.Group("/api/v1")
	groupAPI.POST("company", createCompanyHandler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/company", &jsonRequest)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	createCompanyUseCase.AssertNumberOfCalls(t, "Execute", 1)

}

func Test_CreateCompanyFailsBecauseOfUsecase(t *testing.T) {

	createCompanyUseCase := mocks.CreateCompany{}
	createCompanyHandler := CreateCompany{CreateCompanyUsecase: &createCompanyUseCase}

	request := getCreateCompanyRequestMock()
	jsonRequest := MarshallRequest(request)

	createCompanyUseCase.On("Execute", mock.Anything, &request).Return(errors.New("a random error"))
	r := gin.Default()
	groupAPI := r.Group("/api/v1")
	groupAPI.POST("company", createCompanyHandler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/company", &jsonRequest)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	createCompanyUseCase.AssertNumberOfCalls(t, "Execute", 1)

}

func Test_CreateCompanyFailsBecauseOfBadRequest(t *testing.T) {

	createCompanyUseCase := mocks.CreateCompany{}
	createCompanyHandler := CreateCompany{CreateCompanyUsecase: &createCompanyUseCase}

	request := getCreateCompanyRequestMock()
	request.Name = ""
	jsonRequest := MarshallRequest(request)

	createCompanyUseCase.On("Execute", mock.Anything, &request).Return(errors.New("a random error"))
	r := gin.Default()
	groupAPI := r.Group("/api/v1")
	groupAPI.POST("company", createCompanyHandler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/company", &jsonRequest)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	createCompanyUseCase.AssertNumberOfCalls(t, "Execute", 0)

}

func getCreateCompanyRequestMock() contracts.CreateCompanyRequest {
	return contracts.CreateCompanyRequest {
		Name:   "name",
		Description: "desciption",
		CountryName: "country",
	}
}

func MarshallRequest(request contracts.CreateCompanyRequest) bytes.Buffer {
	requestBodyBytes := new(bytes.Buffer)
	_ = json.NewEncoder(requestBodyBytes).Encode(request)
	return *requestBodyBytes
}
