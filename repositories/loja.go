package repositories

import (
	"api-rest/models"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

// InsertLoja insere uma nova loja
func InsertLoja(l *models.Loja) error {
	Connection()
	defer db.Close()

	query := `INSERT INTO Loja (Nome, RazaoSocial, Endereco, Estado, Cidade, Cep, NumLoja, IdEstabelecimento) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err = db.Exec(query, l.Nome, l.RazaoSocial, l.Endereco, l.Estado, l.Cidade, l.Cep, l.NumLoja, l.IdEstabelecimento)

	if err != nil {
		fmt.Println("Erro ao cadastrar loja:", err.Error())

		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return errors.New("número da loja existente")
		}

		if strings.Contains(err.Error(), "insert or update on table \"store\" violates foreign key constraint") {
			return errors.New("estabelecimento não registrado")
		}

		return err
	}

	fmt.Println("Loja cadastrada com sucesso")
	SelectLoja(l, l.NumLoja)

	return nil
}

// SelectLoja seleciona uma loja
func SelectLoja(l *models.Loja, NumLoja int) (*models.Loja, error) {
	Connection()
	defer db.Close()

	query := `SELECT NumLoja FROM Loja WHERE NumLoja = $1`
	row := db.QueryRow(query, NumLoja)

	err := row.Scan(&l.Id, &l.Nome, &l.RazaoSocial, &l.Endereco, &l.Estado, &l.Cidade, &l.Cep, &l.NumLoja, &l.IdEstabelecimento)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("loja não encontrada")
		}
		fmt.Println("Erro ao listar loja:", err.Error())

		return nil, err
	}

	return l, nil
}

// UpdateLoja atualiza uma loja
func UpdateLoja(l *models.Loja, NumLoja int) error {
	Connection()
	defer db.Close()

	query := "UPDATE Loja SET "
	var args []interface{}

	if l.Nome != "" {
		query += "Nome = $1, "
		args = append(args, l.Nome)
	}
	if l.RazaoSocial != "" {
		query += "RazaoSocial = $2, "
		args = append(args, l.RazaoSocial)
	}
	if l.Endereco != "" {
		query += "Endereco = $3, "
		args = append(args, l.Endereco)
	}
	if l.Estado != "" {
		query += "Estado = $4, "
		args = append(args, l.Estado)
	}
	if l.Cidade != "" {
		query += "Cidade = $5, "
		args = append(args, l.Cidade)
	}
	if l.Cep != "" {
		query += "Cep = $6, "
		args = append(args, l.Cep)
	}
	if l.NumLoja != 0 {
		query += `NumLoja = $7, `
		args = append(args, l.NumLoja)
	}

	query = query[:len(query)-2]

	query += " WHERE NumLoja = $8"
	args = append(args, NumLoja)

	_, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println("Erro ao atualizar a loja:", err.Error())
		return err
	}

	fmt.Println("Loja atualizada com sucesso")
	SelectLoja(l, NumLoja)

	return nil
}

// SelectAllLojas seleciona todas as lojas
func SelectAllLojas(l *models.Loja) ([]models.Loja, error) {
	Connection()
	defer db.Close()

	var lojas []models.Loja

	rows, err := db.Query(`SELECT * FROM Loja`)
	if err != nil {
		fmt.Println("Erro ao listar todas as lojas:", err.Error())

		if strings.Contains(err.Error(), "invalid memory address or nil pointer dereference goroutine") {
			return lojas, errors.New("erro ao listar todas as lojas")
		}

		return lojas, err
	}

	for rows.Next() {
		err := rows.Scan(&l.Id, &l.Nome, &l.RazaoSocial, &l.Endereco, &l.Estado, &l.Cidade, &l.Cep, &l.NumLoja, &l.IdEstabelecimento)
		if err != nil {
			fmt.Println("Erro ao listar todas as lojas:", err.Error())

			return lojas, err
		}

		lojas = append(lojas, *l)
	}
	return lojas, nil
}

// DeleteLoja deleta uma loja
func DeleteLoja(NumLoja int) error {
	Connection()
	defer db.Close()

	query := `DELETE FROM Loja WHERE NumLoja = $1`
	_, err = db.Exec(query, NumLoja)

	if err != nil {
		fmt.Println("Erro ao excluir a loja:", err.Error())

		if strings.Contains(err.Error(), "invalid memory address or nil pointer dereference goroutine") {
			return errors.New("o número da loja inserido é inválido")
		}

		return err
	}

	fmt.Println("Loja excluída com sucesso")
	SelectLoja(&models.Loja{}, NumLoja)

	return nil
}
