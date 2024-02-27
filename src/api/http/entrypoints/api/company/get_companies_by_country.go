package api

import (
	"net/http"

	"github.com/Lucasma95/golang-company-api/src/api/core/entities"
	usecases "github.com/Lucasma95/golang-company-api/src/api/core/usecases/company"
	"github.com/gin-gonic/gin"
)

type GetCompaniesByCountry struct {
	GetCompaniesByCountryUsecase usecases.GetCompaniesByCountry
}

// @Summary      Get company
// @Description  This endpoint allows you to get a company by id
// @Tags         Company
// @Param        command path string true "company_id"
// @Success      200  {object}  []entities.CompanyDTO
// @Failure      500  {object}  error
// @Router       /api/v1/company/country/{country} [get]
func (handler *GetCompaniesByCountry) Handle(c *gin.Context) {

	country := c.Param("country")

	companies, err := handler.GetCompaniesByCountryUsecase.Execute(c.Request.Context(), country)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	dtos := make([]entities.CompanyDTO, len(companies))
	for i, company := range companies {
		dtos[i] = entities.NewCompanyDTO(&company)
	}

	c.JSON(http.StatusOK, dtos)
}