package application

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/Lucasma95/golang-company-api/docs/swagger"
	"github.com/Lucasma95/golang-company-api/src/api/application/config"
	"github.com/Lucasma95/golang-company-api/src/api/http/router"
	"github.com/Lucasma95/golang-company-api/src/api/infrastructure/dependencies"
)

// @title golang-company-api
// @version 1.0
// @BasePath /api
// @query.collection.format multi
func Start() {
	env := config.EnvVars()
	handlers := dependencies.Start(env)

	var host string
	if env.Swagger.HostName == "localhost" {
		host = fmt.Sprintf("%s:%v", env.Swagger.HostName, 3000)
	} else {
		host = env.Swagger.HostName
	}

	swagger.SwaggerInfo.Host = host

	server := gin.New()
	server.Use(gin.LoggerWithWriter(gin.DefaultWriter, "/health"), gin.Recovery())

	router.Init(server, handlers)

	err := server.Run(":3000")
	if err != nil {
		panic(err)
	}
}
