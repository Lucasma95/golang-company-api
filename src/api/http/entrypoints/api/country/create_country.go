package api

import (
	"net/http"
	"log"

	usecases "github.com/Lucasma95/golang-company-api/src/api/core/usecases/country"
	contracts "github.com/Lucasma95/golang-company-api/src/api/http/contracts/country"
	"github.com/gin-gonic/gin"
)

type CreateCountry struct {
	CreateCountryUsecase usecases.CreateCountry
}

// @Summary      Create Country
// @Description  This endpoint allows you to create a country
// @Tags         Country
// @Param        request body contracts.CreateCountryRequest true "Request"
// @Success      201
// @Failure      500
// @Router       /country [post]
func (handler *CreateCountry) Handle(c *gin.Context) {

	var request contracts.CreateCountryRequest
	if err := c.BindJSON(&request); err != nil {
		log.Println("error binding request")
		c.Status(http.StatusBadRequest)
		return
	}

	if err := handler.CreateCountryUsecase.Execute(c.Request.Context(), &request); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}
