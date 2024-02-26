package handlers

import (
	"api-rest/models"
	"api-rest/repositories"
	"api-rest/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func NewLoja(c echo.Context) error {
	request := &models.Loja{}
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, MensagemErro(err.Error()))
	}

	if err := services.InsertLoja(request); err != nil {
		return c.JSON(http.StatusBadRequest, MensagemErro(err.Error()))
	}

	return c.JSON(http.StatusOK, MensagemSucesso())
}

func GetLoja(c echo.Context) error {
	NumLoja, err := strconv.Atoi(c.Param("NumLoja"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, MensagemErro("Parâmetro inválido"))
	}

	loja, err := repositories.SelectLoja(&models.Loja{}, NumLoja)
	if err != nil {
		return c.JSON(http.StatusBadRequest, MensagemErro(err.Error()))
	}

	return c.JSON(http.StatusOK, DadosMensagemSucesso(loja))
}

func UpdateLoja(c echo.Context) error {
	request := &models.Loja{}
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, MensagemErro(err.Error()))
	}

	NumLoja, err := strconv.Atoi(c.Param("NumLoja"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, MensagemErro("Parâmetro inválido"))
	}

	if err := services.UpdateLoja(request, NumLoja); err != nil {
		return c.JSON(http.StatusBadRequest, MensagemErro(err.Error()))
	}

	return c.JSON(http.StatusOK, MensagemSucesso())
}

func GetAllLojas(c echo.Context) error {
	lojas, err := repositories.SelectAllLojas(&models.Loja{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, MensagemErro(err.Error()))
	}

	return c.JSON(http.StatusOK, DadosMensagemSucesso(lojas))
}

func DeleteLoja(c echo.Context) error {
	NumLoja, err := strconv.Atoi(c.Param("NumLoja"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, MensagemErro("Parâmetro inválido"))
	}

	if err := repositories.DeleteLoja(NumLoja); err != nil {
		return c.JSON(http.StatusBadRequest, MensagemErro(err.Error()))
	}

	return c.JSON(http.StatusOK, MensagemSucesso())
}
