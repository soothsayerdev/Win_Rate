package models

import (
	"database/sql"
	"fmt"
	"log"
)

// Deck struct defines the properties of a deck
type Deck struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	DeckName string `json:"deck_name"`
}

// CreateDeck function inserts a new deck into the database
func (d *Deck) CreateDeck(db *sql.DB) error {
	query := `
		INSERT INTO decks (user_id, deck_name)
		VALUES ($1, $2) RETURNING id
	`
	err := db.QueryRow(query, d.UserID, d.DeckName).Scan(&d.ID)
	if err != nil {
		return fmt.Errorf("failed to create deck: %w", err)
	}
	log.Printf("Deck created with ID: %d", d.ID)
	return nil
}

// GetDecksByUser function fetches all decks for a specific user
func GetDecksByUser(db *sql.DB, userID int) ([]Deck, error) {
	rows, err := db.Query("SELECT id, user_id, deck_name FROM decks WHERE user_id = $1", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch decks for user %d: %w", userID, err)
	}
	defer rows.Close()

	var decks []Deck
	for rows.Next() {
		var deck Deck
		if err := rows.Scan(&deck.ID, &deck.UserID, &deck.DeckName); err != nil {
			return nil, fmt.Errorf("failed to scan deck: %w", err)
		}
		decks = append(decks, deck)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error encountered while iterating rows: %w", err)
	}

	return decks, nil
}

// GetDeckByID function retrieves a deck by its ID
func GetDeckByID(db *sql.DB, deckID int) (*Deck, error) {
	var deck Deck
	query := "SELECT id, user_id, deck_name FROM decks WHERE id = $1"
	err := db.QueryRow(query, deckID).Scan(&deck.ID, &deck.UserID, &deck.DeckName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("deck with ID %d not found", deckID)
		}
		return nil, fmt.Errorf("failed to fetch deck with ID %d: %w", deckID, err)
	}
	return &deck, nil
}

// UpdateDeck function updates an existing deck's name
func (d *Deck) UpdateDeck(db *sql.DB) error {
	query := "UPDATE decks SET deck_name = $1 WHERE id = $2 AND user_id = $3"
	_, err := db.Exec(query, d.DeckName, d.ID, d.UserID)
	if err != nil {
		return fmt.Errorf("failed to update deck: %w", err)
	}
	log.Printf("Deck with ID %d updated", d.ID)
	return nil
}

// DeleteDeck function removes a deck from the database
func (d *Deck) DeleteDeck(db *sql.DB) error {
	query := "DELETE FROM decks WHERE id = $1 AND user_id = $2"
	_, err := db.Exec(query, d.ID, d.UserID)
	if err != nil {
		return fmt.Errorf("failed to delete deck: %w", err)
	}
	log.Printf("Deck with ID %d deleted", d.ID)
	return nil
}
