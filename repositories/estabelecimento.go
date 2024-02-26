package repositories

import (
	"api-rest/models"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

// InsertEstabelecimento adiciona um novo estabelecimento
func InsertEstabelecimento(e *models.Estabelecimento) error {
	Connection()
	defer db.Close()

	query := `INSERT INTO Estabelecimento (Nome, RazaoSocial, Endereco, Estado, Cidade, Cep, NumEstabelecimento) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := db.Exec(query, e.Nome, e.RazaoSocial, e.Endereco, e.Estado, e.Cidade, e.Cep, e.NumEstabelecimento)

	if err != nil {
		fmt.Println("Erro ao cadastrar estabelecimento:", err.Error())

		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return errors.New("número de estabelecimento existente")
		}

		return err
	}

	fmt.Println("Estabelecimento cadastrado com sucesso")
	SelectEstabelecimento(e, e.NumEstabelecimento)

	return nil
}

// SelectEstabelecimento seleciona um estabelecimento
func SelectEstabelecimento(e *models.Estabelecimento, NumEstabelecimento int) (*models.Estabelecimento, error) {
	Connection()
	defer db.Close()

	query := `SELECT * FROM Estabelecimento WHERE NumEstabelecimento = $1`
	row := db.QueryRow(query, NumEstabelecimento)

	err := row.Scan(&e.Id, &e.Nome, &e.RazaoSocial, &e.Endereco, &e.Estado, &e.Cidade, &e.Cep, &e.NumEstabelecimento)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("estabelecimento não encontrado")
		}
		fmt.Println("Erro ao listar estabelecimento:", err.Error())

		return nil, err
	}

	return e, nil
}

// UpdateEstabelecimento atualiza um estabelecimento
func UpdateEstabelecimento(e *models.Estabelecimento, NumEstabelecimento int) error {
	Connection()
	defer db.Close()

	query := "UPDATE Estabelecimento SET "
	var args []interface{}

	if e.Nome != "" {
		query += "Nome = $1, "
		args = append(args, e.Nome)
	}
	if e.RazaoSocial != "" {
		query += "RazaoSocial = $2, "
		args = append(args, e.RazaoSocial)
	}
	if e.Endereco != "" {
		query += "Endereco = $3, "
		args = append(args, e.Endereco)
	}
	if e.Estado != "" {
		query += "Estado = $4, "
		args = append(args, e.Estado)
	}
	if e.Cidade != "" {
		query += "Cidade = $5, "
		args = append(args, e.Cidade)
	}
	if e.Cep != "" {
		query += "Cep = $6, "
		args = append(args, e.Cep)
	}

	query = query[:len(query)-2]

	query += " WHERE NumEstabelecimento = $7"
	args = append(args, NumEstabelecimento)

	_, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println("Erro ao atualizar o estabelecimento:", err.Error())
		return err
	}

	fmt.Println("Estabelecimento atualizado com sucesso")
	SelectEstabelecimento(e, NumEstabelecimento)

	return nil
}

// SelectAllEstabelecimentos seleciona todos os estabelecimentos
func SelectAllEstabelecimentos(e *models.Estabelecimento) ([]models.Estabelecimento, error) {
	Connection()
	defer db.Close()

	var estabelecimentos []models.Estabelecimento

	rows, err := db.Query(`SELECT * FROM Estabelecimento`)
	if err != nil {
		fmt.Println("Erro ao listar todos os estabelecimentos:", err.Error())

		if strings.Contains(err.Error(), "invalid memory address or nil pointer dereference goroutine") {
			return estabelecimentos, errors.New("erro ao listar todos os estabelecimentos")
		}

		return estabelecimentos, err
	}

	for rows.Next() {
		err := rows.Scan(&e.Id, &e.Nome, &e.RazaoSocial, &e.Endereco, &e.Estado, &e.Cidade, &e.Cep, &e.NumEstabelecimento)
		if err != nil {
			fmt.Println("Erro ao listar todos os estabelecimentos:", err.Error())

			return estabelecimentos, err
		}

		estabelecimentos = append(estabelecimentos, *e)
	}

	return estabelecimentos, nil
}

// DeleteEstabelecimento deleta um estabelecimento
func DeleteEstabelecimento(NumEstabelecimento int) error {
	Connection()
	defer db.Close()

	query := `DELETE FROM Estabelecimento WHERE NumEstabelecimento = $1`
	_, err = db.Exec(query, NumEstabelecimento)

	if err != nil {
		fmt.Println("Erro ao excluir o estabelecimento:", err.Error())

		if strings.Contains(err.Error(), "invalid memory address or nil pointer dereference goroutine") {
			return errors.New("o número do estabelecimento inserido é inválido")
		}

		return err
	}

	fmt.Println("Estabelecimento excluído com sucesso")
	SelectEstabelecimento(&models.Estabelecimento{}, NumEstabelecimento)

	return nil
}
