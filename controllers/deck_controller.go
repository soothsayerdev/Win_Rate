package controller

import (
	"database/sql"
	"fmt"
	"log"
)

// Função para criar um novo deck para o usuário
func createDeck(db *sql.DB, userID int, deckName string) error {
	// Validação de entrada
	if userID <= 0 {
		return fmt.Errorf("userID inválido: %d", userID)
	}
	if deckName == "" {
		return fmt.Errorf("deckName não pode ser vazio")
	}

	// Verificando se o usuário existe no banco antes de criar o deck
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`
	err := db.QueryRow(query, userID).Scan(&exists)
	if err != nil {
		log.Printf("Erro ao verificar usuário: %v", err)
		return fmt.Errorf("erro ao verificar se o usuário existe: %w", err)
	}

	if !exists {
		return fmt.Errorf("usuário com ID %d não encontrado", userID)
	}

	// Usando transação para garantir a consistência
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Erro ao iniciar transação: %v", err)
		return fmt.Errorf("erro ao iniciar transação: %w", err)
	}

	// Commit ou rollback no final
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// Inserção no banco de dados
	insertQuery := `
		INSERT INTO decks (user_id, deck_name)
		VALUES ($1, $2)
	`
	_, err = tx.Exec(insertQuery, userID, deckName)
	if err != nil {
		log.Printf("Erro ao inserir deck: %v", err)
		return fmt.Errorf("erro ao inserir deck: %w", err)
	}

	return nil
}
