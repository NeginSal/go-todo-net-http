package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NeginSal/go-todo-net-http/db"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Username == "" || user.Password == "" {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	if err != nil {
		http.Error(w, "Error during registration. Perhaps the username is already taken.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "sign up is successful")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "just POST request is allowed!", http.StatusMethodNotAllowed)
		return
	}

	var input User
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.Username == "" || input.Password == "" {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	var storedPassword string
	err = db.DB.QueryRow("SELECT password FROM users WHERE username = ?", input.Username).Scan(&storedPassword)
	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	if storedPassword != input.Password {
		http.Error(w, "Wrong password", http.StatusUnauthorized)
		return
	}
	fmt.Fprintln(w, "login is successful")
}
