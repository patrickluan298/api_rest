package repositories

import (
	"api-rest/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertEstabelecimento(t *testing.T) {
	test := &models.Estabelecimento{
		Nome:               "Shopping Rio Major",
		RazaoSocial:        "NexTech",
		Endereco:           "Avenida Espírito Santo",
		Estado:             "Paraíba",
		Cidade:             "João Pessoa",
		Cep:                "11223344",
		NumEstabelecimento: 1676,
	}

	err := InsertEstabelecimento(test)
	assert.NoError(t, err)
}

func TestSelectEstabelecimento(t *testing.T) {
	estabelecimento, err := SelectEstabelecimento(&models.Estabelecimento{}, 1671)

	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	if estabelecimento == nil {
		t.Error("O estabelecimento não deveria ser nulo")
	}
}

func TestUpdateEstabelecimento(t *testing.T) {
	test := &models.Estabelecimento{
		Nome:        "Padaria do João",
		RazaoSocial: "Padaria do João Ltda.",
		Endereco:    "Rua das Flores, 123",
		Estado:      "SP",
		Cidade:      "São Paulo",
		Cep:         "01234-567",
	}

	err := UpdateEstabelecimento(test, 1699)
	assert.NoError(t, err)
}

func TestSelectAllEstabelecimentos(t *testing.T) {
	estabelecimentos, err := SelectAllEstabelecimentos(&models.Estabelecimento{})

	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	if estabelecimentos == nil {
		t.Error("O estabelecimento não deveria ser nulo")
	}
}

func TestDeleteEstabelecimento(t *testing.T) {
	err := DeleteEstabelecimento(1)
	assert.NoError(t, err)
}
