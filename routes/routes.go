package routes

import (
	"api-rest/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// HandleRequest responsável pelas rotas da aplicação
func HandleRequest(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", handlers.HealthCheck)

	e.POST("/api/v1/estabelecimentos", handlers.NewEstabelecimento)
	e.GET("/api/v1/estabelecimentos", handlers.GetAllEstabelecimentos)
	e.GET("/api/v1/estabelecimentos/:NumEstabelecimento", handlers.GetEstabelecimento)
	e.PUT("/api/v1/estabelecimentos/:NumEstabelecimento", handlers.UpdateEstabelecimento)
	e.DELETE("/api/v1/estabelecimentos/:NumEstabelecimento", handlers.DeleteEstabelecimento)

	e.POST("/api/v1/lojas", handlers.NewLoja)
	e.GET("/api/v1/lojas", handlers.GetAllLojas)
	e.GET("/api/v1/lojas/:NumLoja", handlers.GetLoja)
	e.PUT("/api/v1/lojas/:NumLoja", handlers.UpdateLoja)
	e.DELETE("/api/v1/lojas/:NumLoja", handlers.DeleteLoja)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	handlers.StartServer(e)
}
