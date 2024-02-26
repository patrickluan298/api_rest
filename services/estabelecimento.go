package services

import (
	"api-rest/models"
	"api-rest/repositories"
	"fmt"

	"gopkg.in/validator.v2"
)

// ValidateEstabelecimento valida a estrutura do estabelecimento
func ValidateEstabelecimento(estabelecimento *models.Estabelecimento) error {
	if err := validator.Validate(estabelecimento); err != nil {
		return fmt.Errorf("erro na validação do estabelecimento: %v", err)
	}
	
	return nil
}

// InsertEstabelecimento insere um novo estabelecimento
func InsertEstabelecimento(estabelecimento *models.Estabelecimento) error {
	if err := ValidateEstabelecimento(estabelecimento); err != nil {
		return fmt.Errorf("erro ao validar estabelecimento: %v", err)
	}

	if err := repositories.InsertEstabelecimento(estabelecimento); err != nil {
		return fmt.Errorf("erro ao inserir estabelecimento no banco de dados: %v", err)
	}

	return nil
}

// UpdateEstabelecimento atualiza um estabelecimento existente
func UpdateEstabelecimento(estabelecimento *models.Estabelecimento, id int) error {
	if err := ValidateEstabelecimento(estabelecimento); err != nil {
		return fmt.Errorf("erro ao validar estabelecimento: %v", err)
	}

	err := repositories.UpdateEstabelecimento(estabelecimento, id)
	if err != nil {
		return fmt.Errorf("erro ao atualizar estabelecimento no banco de dados: %v", err)
	}

	return nil
}
