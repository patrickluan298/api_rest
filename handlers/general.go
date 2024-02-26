package handlers

import (
	"fmt"
	"net/http"

	"api-rest/models"

	"github.com/labstack/echo/v4"
)

func HealthCheck(c echo.Context) error {
	message := models.Mensagem{
		Mensagem: "Aplicação up e em pleno funcionamento",
	}
	return c.JSON(http.StatusOK, message)
}

func MensagemErro(message string) interface{} {
	return models.Mensagem{
		Mensagem: message,
	}
}

func MensagemSucesso() interface{} {
	return models.Mensagem{
		Mensagem: "OK",
	}
}

func DadosMensagemSucesso(dados interface{}) interface{} {
	return models.MensagemDeDados{
		Mensagem: "OK",
		Dados:    dados,
	}
}

func StartServer(e *echo.Echo) {
	fmt.Println("Swagger docs at: http://localhost:1323/swagger/")
	e.Logger.Fatal(e.Start(":1323"))
}
