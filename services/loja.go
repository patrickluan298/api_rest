package services

import (
	"api-rest/models"
	"api-rest/repositories"
	"fmt"

	"gopkg.in/validator.v2"
)

// ValidateLoja valida a estrutura da loja
func ValidateLoja(loja *models.Loja) error {
	if err := validator.Validate(loja); err != nil {
		return fmt.Errorf("dados da loja estão inválidos: %v", err)
	}

	return nil
}

// InsertLoja insere uma nova loja no banco de dados
func InsertLoja(request *models.Loja) error {
	if err := ValidateLoja(request); err != nil {
		return fmt.Errorf("erro ao validar dados da loja: %v", err)
	}

	if err := repositories.InsertLoja(request); err != nil {
		return fmt.Errorf("erro ao inserir loja no banco de dados: %v", err)
	}

	return nil
}

// UpdateLoja valida e atualiza uma loja existente no banco de dados
func UpdateLoja(request *models.Loja, id int) error {
	if err := ValidateLoja(request); err != nil {
		return fmt.Errorf("erro ao validar dados da loja: %v", err)
	}

	if err := repositories.UpdateLoja(request, id); err != nil {
		return fmt.Errorf("erro ao atualizar loja no banco de dados: %v", err)
	}

	return nil
}
