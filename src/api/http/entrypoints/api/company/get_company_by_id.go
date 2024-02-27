package api

import (
	"net/http"

	"github.com/Lucasma95/golang-company-api/src/api/core/entities"
	usecases "github.com/Lucasma95/golang-company-api/src/api/core/usecases/company"
	"github.com/gin-gonic/gin"
)

type GetCompanyByID struct {
	GetCompanyByIDUsecase usecases.GetCompanyByID
}

// @Summary      Get company
// @Description  This endpoint allows you to get a company by id
// @Tags         Company
// @Param        command path string true "company_id"
// @Success      200  {object}  entities.Company
// @Failure      500  {object}  error
// @Router       /subscription/config/{company_id} [get]
func (handler *GetCompanyByID) Handle(c *gin.Context) {

	compnayID := c.Param("company_id")

	company, err := handler.GetCompanyByIDUsecase.Execute(c.Request.Context(), compnayID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	} else if company == nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, entities.NewCompanyDTO(company))
}
