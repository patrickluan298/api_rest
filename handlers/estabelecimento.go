package handlers

import (
	"api-rest/models"
	"api-rest/repositories"
	"api-rest/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func NewEstabelecimento(c echo.Context) error {
	request := &models.Estabelecimento{}
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, MensagemErro(err.Error()))
	}

	if err := services.InsertEstabelecimento(request); err != nil {
		return c.JSON(http.StatusBadRequest, MensagemErro(err.Error()))
	}

	return c.JSON(http.StatusOK, MensagemSucesso())
}

func GetEstabelecimento(c echo.Context) error {
	NumEstabelecimento, err := strconv.Atoi(c.Param("NumEstabelecimento"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, MensagemErro("Parâmetro inválido"))
	}

	estabelecimento, err := repositories.SelectEstabelecimento(&models.Estabelecimento{}, NumEstabelecimento)
	if err != nil {
		return c.JSON(http.StatusBadRequest, MensagemErro(err.Error()))
	}

	return c.JSON(http.StatusOK, DadosMensagemSucesso(estabelecimento))
}

func UpdateEstabelecimento(c echo.Context) error {
	request := &models.Estabelecimento{}
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, MensagemErro(err.Error()))
	}

	NumEstabelecimento, err := strconv.Atoi(c.Param("NumEstabelecimento"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, MensagemErro("Parâmetro inválido"))
	}

	if err := services.UpdateEstabelecimento(request, NumEstabelecimento); err != nil {
		return c.JSON(http.StatusBadRequest, MensagemErro(err.Error()))
	}

	return c.JSON(http.StatusOK, MensagemSucesso())
}

func GetAllEstabelecimentos(c echo.Context) error {
	estabelecimentos, err := repositories.SelectAllEstabelecimentos(&models.Estabelecimento{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, MensagemErro(err.Error()))
	}

	return c.JSON(http.StatusOK, DadosMensagemSucesso(estabelecimentos))
}

func DeleteEstabelecimento(c echo.Context) error {
	NumEstabelecimento, err := strconv.Atoi(c.Param("NumEstabelecimento"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, MensagemErro("Parâmetro inválido"))
	}

	if err := repositories.DeleteEstabelecimento(NumEstabelecimento); err != nil {
		return c.JSON(http.StatusBadRequest, MensagemErro(err.Error()))
	}

	return c.JSON(http.StatusOK, MensagemSucesso())
}
