package api

import (
	"bytes"
	"errors"
	"os"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	contracts "github.com/Lucasma95/golang-company-api/src/api/http/contracts/country"
	mocks "github.com/Lucasma95/golang-company-api/src/api/test/mocks/usecases"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func init() {
	_ = os.Setenv("GIN_MODE", "test")
}

func Test_CreateCountrySuccessfully(t *testing.T) {

	createCountryUseCase := mocks.CreateCountry{}
	createCountryHandler := CreateCountry{CreateCountryUsecase: &createCountryUseCase}

	request := getCreateCountryRequestMock()
	jsonRequest := MarshallRequest(request)

	createCountryUseCase.On("Execute", mock.Anything, &request).Return(nil)
	r := gin.Default()
	groupAPI := r.Group("/api/v1")
	groupAPI.POST("country", createCountryHandler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/country", &jsonRequest)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	createCountryUseCase.AssertNumberOfCalls(t, "Execute", 1)

}

func Test_CreateCountryFailsBecauseOfUsecase(t *testing.T) {

	createCountryUseCase := mocks.CreateCountry{}
	createCountryHandler := CreateCountry{CreateCountryUsecase: &createCountryUseCase}

	request := getCreateCountryRequestMock()
	jsonRequest := MarshallRequest(request)

	createCountryUseCase.On("Execute", mock.Anything, &request).Return(errors.New("a random error"))
	r := gin.Default()
	groupAPI := r.Group("/api/v1")
	groupAPI.POST("country", createCountryHandler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/country", &jsonRequest)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	createCountryUseCase.AssertNumberOfCalls(t, "Execute", 1)

}

func Test_CreateCountryFailsBecauseOfBadRequest(t *testing.T) {

	createCountryUseCase := mocks.CreateCountry{}
	createCountryHandler := CreateCountry{CreateCountryUsecase: &createCountryUseCase}

	request := getCreateCountryRequestMock()
	request.Name = ""
	jsonRequest := MarshallRequest(request)
	
	r := gin.Default()
	groupAPI := r.Group("/api/v1")
	groupAPI.POST("country", createCountryHandler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/country", &jsonRequest)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	createCountryUseCase.AssertNumberOfCalls(t, "Execute", 0)

}

func getCreateCountryRequestMock() contracts.CreateCountryRequest {
	return contracts.CreateCountryRequest{
		Name:      "country",
		Continent: "continent",
	}
}

func MarshallRequest(request contracts.CreateCountryRequest) bytes.Buffer {
	requestBodyBytes := new(bytes.Buffer)
	_ = json.NewEncoder(requestBodyBytes).Encode(request)
	return *requestBodyBytes
}
