package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./todo.db")
	if err != nil {
		panic(err)
	}
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`

	_, err = DB.Exec(createUserTable)
	if err != nil {
		panic(err)
	}

	createTodoTable := `
  CREATE TABLE IF NOT EXISTS todos (
	  id INTEGER PRIMARY KEY AUTOINCREMENT,
	  title TEXT NOT NULL,
	  completed BOOLEAN NOT NULL DEFAULT 0,
	  user_id INTEGER NOT NULL,
	  FOREIGN KEY(user_id) REFERENCES users(id)
  );`

	_, err = DB.Exec(createTodoTable)
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Connected to SQLite and ensured user table exists")
}
