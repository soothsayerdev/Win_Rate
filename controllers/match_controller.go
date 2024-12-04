package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"database/sql"
)

// Definir a estrutura do "match" conforme o banco de dados
type Match struct {
	UserDeckID      int `json:"userDeckID"`
	OpponentDeckID  int `json:"opponentDeckID"`
	Victories       int `json:"victories"`
	Defeats         int `json:"defeats"`
}

// Função para calcular o win rate
func calculateWinRate(victories, defeats int) float64 {
	if victories+defeats == 0 {
		return 0.0
	}
	return float64(victories) / float64(victories+defeats)
}

// Função para criar uma nova partida
func createMatch(w http.ResponseWriter, r *http.Request) {
	// Definir a variável para armazenar o corpo da requisição
	var match Match

	// Decodificar o JSON recebido
	err := json.NewDecoder(r.Body).Decode(&match)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao decodificar dados: %v", err), http.StatusBadRequest)
		return
	}

	// Validar dados obrigatórios
	if match.UserDeckID == 0 || match.OpponentDeckID == 0 {
		http.Error(w, "Os IDs dos decks devem ser válidos.", http.StatusBadRequest)
		return
	}

	// Calcular a taxa de vitória
	winRate := calculateWinRate(match.Victories, match.Defeats)

	// Inserir a partida no banco de dados
	_, err = db.Exec("INSERT INTO matches (user_deck_id, opponent_deck_id, victories, defeats) VALUES (?, ?, ?, ?)",
		match.UserDeckID, match.OpponentDeckID, match.Victories, match.Defeats)
	if err != nil {
		log.Printf("Erro ao criar a partida: %v", err)
		http.Error(w, "Erro ao criar partida", http.StatusInternalServerError)
		return
	}

	// Criar a resposta
	response := map[string]interface{}{
		"message":       "Partida criada com sucesso!",
		"winRate":       winRate,
		"userDeckID":    match.UserDeckID,
		"opponentDeckID": match.OpponentDeckID,
	}

	// Definir o status da resposta e enviar o JSON
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Erro ao enviar a resposta: %v", err)
		http.Error(w, "Erro ao enviar resposta", http.StatusInternalServerError)
	}
}
