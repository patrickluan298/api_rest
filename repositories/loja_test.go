package repositories

import (
	"api-rest/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertLoja(t *testing.T) {
	test := &models.Loja{
		Nome:              "Loja de teste",
		RazaoSocial:       "Teste LTDA",
		Endereco:          "Rua de teste, 123",
		Estado:            "SP",
		Cidade:            "São Paulo",
		Cep:               "12345-678",
		NumLoja:           400,
		IdEstabelecimento: 1671,
	}

	err := InsertLoja(test)
	assert.NoError(t, err)
}

func TestSelectLoja(t *testing.T) {
	loja, err := SelectLoja(&models.Loja{}, 400)

	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	if loja == nil {
		t.Error("A loja não deveria ser nula")
	}
}

func TestUpdateLoja(t *testing.T) {
	test := &models.Loja{
		Nome:              "Loja de teste",
		RazaoSocial:       "Teste LTDA",
		Endereco:          "Rua de teste, 123",
		Estado:            "SP",
		Cidade:            "Rio de Janeiro",
		Cep:               "12345-678",
		NumLoja:           1,
		IdEstabelecimento: 1,
	}

	err := UpdateLoja(test, 1)
	assert.NoError(t, err)
}

func TestSelectAllLojas(t *testing.T) {
	lojas, err := SelectAllLojas(&models.Loja{})

	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	if lojas == nil {
		t.Error("A loja não deveria ser nula")
	}
}

func TestDeleteLoja(t *testing.T) {
	err := DeleteLoja(1)
	assert.NoError(t, err)
}
