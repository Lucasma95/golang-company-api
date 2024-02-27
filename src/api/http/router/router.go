package router

import (
	"github.com/Lucasma95/golang-company-api/src/api/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(server *gin.Engine, handlers *dependencies.HandlerContainer) {
	registerHealthRoute(server, handlers)
	registerSwaggerRoute(server)
	registerAPIRouter(server, handlers)
}

func registerHealthRoute(server *gin.Engine, handlers *dependencies.HandlerContainer) {
	server.GET("/health", handlers.Health.Handle)
}

func registerSwaggerRoute(server *gin.Engine) {
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func registerAPIRouter(server *gin.Engine, handlers *dependencies.HandlerContainer) {
	groupAPI := server.Group("/api/v1")

	groupAPI.POST("country", handlers.CreateCountry.Handle)

	groupAPI.POST("company", handlers.CreateCompany.Handle)
	groupAPI.GET("company/:id", handlers.GetCompanyByID.Handle)
	groupAPI.GET("company/country/:country_name", handlers.GetCompanyByCountry.Handle)
	groupAPI.GET("company/continent/:continent", handlers.GetCompanyByContinent.Handle)
}
