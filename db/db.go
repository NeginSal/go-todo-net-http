package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./todo.db")
	if err != nil {
		log.Fatal("❌ Error connecting to the database:", err)
	}

	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`

	_, err = DB.Exec(createUserTable)
	if err != nil {
		log.Fatal("❌ Error creating users table:", err)
	}

	fmt.Println("✅ SQLite database connected and users table created.")
}
