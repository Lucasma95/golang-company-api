package api

import (
	"log"
	"net/http"

	usecases "github.com/Lucasma95/golang-company-api/src/api/core/usecases/company"
	contracts "github.com/Lucasma95/golang-company-api/src/api/http/contracts/company"
	"github.com/gin-gonic/gin"
)

type CreateCompany struct {
	CreateCompanyUsecase usecases.CreateCompany
}

// @Summary      Create Company
// @Description  This endpoint allows you to create a company
// @Tags         Company
// @Param        request body contracts.CreateCompanyRequest true "Request"
// @Success      201
// @Failure      500
// @Router       /company [post]
func (handler *CreateCompany) Handle(c *gin.Context) {

	var request contracts.CreateCompanyRequest
	if err := c.BindJSON(&request); err != nil {
		log.Println("error binding request")
		c.Status(http.StatusBadRequest)
		return
	}

	if err := handler.CreateCompanyUsecase.Execute(c.Request.Context(), &request); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}
