package dependencies

import (
	"github.com/Lucasma95/golang-company-api/src/api/application/config"
	companyUsecases "github.com/Lucasma95/golang-company-api/src/api/core/usecases/company"
	countryUsecases "github.com/Lucasma95/golang-company-api/src/api/core/usecases/country"
	"github.com/Lucasma95/golang-company-api/src/api/http/entrypoints"
	companyHandlers "github.com/Lucasma95/golang-company-api/src/api/http/entrypoints/api/company"
	countryHandlers "github.com/Lucasma95/golang-company-api/src/api/http/entrypoints/api/country"
	database "github.com/Lucasma95/golang-company-api/src/api/infrastructure/repositories/database"
	dbClient "github.com/Lucasma95/golang-company-api/src/api/infrastructure/repositories/database/client"
)

type HandlerContainer struct {
	GetCompanyByCountry   entrypoints.Handler
	GetCompanyByContinent entrypoints.Handler
	GetCompanyByID        entrypoints.Handler
	CreateCompany         entrypoints.Handler
	CreateCountry         entrypoints.Handler
	Health                entrypoints.Handler
}

func Start(env config.Config) *HandlerContainer {

	var handlers HandlerContainer

	//clients
	dbClient := dbClient.New(&env)

	//repositories
	countryRepository := database.CountryRepository{Client: dbClient}
	companyRepository := database.CompanyRepository{Client: dbClient}

	//company usecases
	getCompanyByID := companyUsecases.GetCompanyByIDImpl{
		CompanyRepository: &companyRepository,
	}
	getCompaniesByCountry := companyUsecases.GetCompaniesByCountryImpl{
		CompanyRepository: &companyRepository,
	}
	getCompaniesByContinent := companyUsecases.GetCompaniesByContinentImpl{
		CompanyRepository: &companyRepository,
	}
	createCompany := companyUsecases.CreateCompanyImpl{
		CompanyRepository: &companyRepository,
	}

	//country usecases
	createCountry := countryUsecases.CreateCountryImpl{
		CountryRepository: &countryRepository,
	}

	//company handlers
	handlers.CreateCompany = &companyHandlers.CreateCompany{CreateCompanyUsecase: &createCompany}
	handlers.GetCompanyByID = &companyHandlers.GetCompanyByID{GetCompanyByIDUsecase: &getCompanyByID}
	handlers.GetCompanyByCountry = &companyHandlers.GetCompaniesByCountry{GetCompaniesByCountryUsecase: &getCompaniesByCountry}
	handlers.GetCompanyByContinent = &companyHandlers.GetCompaniesByContinent{GetCompaniesByContinentUsecase: &getCompaniesByContinent}

	//Country Handlers
	handlers.CreateCountry = &countryHandlers.CreateCountry{CreateCountryUsecase: &createCountry}

	return &handlers
}
