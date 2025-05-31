package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NeginSal/go-todo-net-http/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
        return
    }

    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil || user.Username == "" || user.Password == "" {
        http.Error(w, "Invalid data", http.StatusBadRequest)
        return
    }

    // Hashing the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Error encrypting", http.StatusInternalServerError)
        return
    }

    _, err = db.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, string(hashedPassword))
    if err != nil {
        http.Error(w, "Error during registration. Username might be duplicate.", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    fmt.Fprintln(w, "✅ Registration successful (password hashed)")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
        return
    }

    var input User
    err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil || input.Username == "" || input.Password == "" {
        http.Error(w, "Invalid data", http.StatusBadRequest)
        return
    }

    var storedHashedPassword string
    err = db.DB.QueryRow("SELECT password FROM users WHERE username = ?", input.Username).Scan(&storedHashedPassword)
    if err == sql.ErrNoRows {
        http.Error(w, "User not found", http.StatusUnauthorized)
        return
    } else if err != nil {
        http.Error(w, "Server error", http.StatusInternalServerError)
        return
    }

    // Comparing the entered password with the stored hash
    err = bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(input.Password))
    if err != nil {
        http.Error(w, "Wrong password", http.StatusUnauthorized)
        return
    }

    fmt.Fprintln(w, "✅ Login successful")
}
