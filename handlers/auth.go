package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NeginSal/go-todo-net-http/database"
	"github.com/NeginSal/go-todo-net-http/models"
	"github.com/NeginSal/go-todo-net-http/utils"
)

// RegisterHandler handles user registration
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Username == "" || user.Password == "" || user.Email == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Failed to encrypt password", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec(
		"INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
		user.Username, user.Email, hashedPassword,
	)
	if err != nil {
		http.Error(w, "User already exists or DB error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "✅ Registration successful")
}

// LoginHandler handles user login and returns a JWT token
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var input models.User
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.Username == "" || input.Password == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var storedHashedPassword string
	err = database.DB.QueryRow("SELECT password FROM users WHERE username = ?", input.Username).Scan(&storedHashedPassword)
	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	if !utils.CheckPasswordHash(input.Password, storedHashedPassword) {
		http.Error(w, "Wrong password", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(input.Username)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "✅ Login successful\nToken: %s", token)
}
