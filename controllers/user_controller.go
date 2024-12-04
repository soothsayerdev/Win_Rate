package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"winrate/models"
	"winrate/utils"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser handles the user registration process.
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	
	// Decodificando os dados da requisição
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		log.Printf("Error decoding user data: %v", err)
		return
	}

	// Validando se os campos necessários estão preenchidos
	if user.Username == "" || user.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// Criptografando a senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error encrypting password", http.StatusInternalServerError)
		log.Printf("Error encrypting password: %v", err)
		return
	}
	user.Password = string(hashedPassword)

	// Registrando o usuário no banco de dados
	err = user.Register(utils.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		log.Printf("Error registering user: %v", err)
		return
	}

	// Retornando a resposta de sucesso
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

// LoginUser handles the user login process.
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	
	// Decodificando os dados da requisição
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		log.Printf("Error decoding user data: %v", err)
		return
	}

	// Validando se os campos de login estão preenchidos
	if user.Username == "" || user.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// Verificando se o usuário existe no banco de dados
	storedUser, err := user.Login(utils.DB)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		log.Printf("Error finding user: %v", err)
		return
	}

	// Comparando as senhas usando bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		log.Printf("Invalid credentials for user %s", user.Username)
		return
	}

	// Gerando um token de autenticação (por exemplo, com JWT)
	token, err := utils.GenerateJWT(storedUser)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		log.Printf("Error generating token: %v", err)
		return
	}

	// Retornando a resposta de sucesso com o token
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login successful",
		"token":   token,
	})
}
