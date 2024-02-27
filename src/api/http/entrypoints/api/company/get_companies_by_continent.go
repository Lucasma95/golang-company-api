package api

import (
	"net/http"

	"github.com/Lucasma95/golang-company-api/src/api/core/entities"
	usecases "github.com/Lucasma95/golang-company-api/src/api/core/usecases/company"
	"github.com/gin-gonic/gin"
)

type GetCompaniesByContinent struct {
	GetCompaniesByContinentUsecase usecases.GetCompaniesByContinent
}

// @Summary      Get company
// @Description  This endpoint allows you to get a company by continent
// @Tags         Company
// @Param        command path string true "continent"
// @Success      200  {object}  []entities.CompanyDTO
// @Failure      500  {object}  error
// @Router       /api/v1/company/continent/{continent} [get]
func (handler *GetCompaniesByContinent) Handle(c *gin.Context) {

	continent := c.Param("continent")

	companies, err := handler.GetCompaniesByContinentUsecase.Execute(c.Request.Context(), continent)
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
