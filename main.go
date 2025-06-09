package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NeginSal/go-todo-net-http/config"
	"github.com/NeginSal/go-todo-net-http/database"
	"github.com/NeginSal/go-todo-net-http/handlers"
)

func main() {
	config.InitEnv()
	database.ConnectDB()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the TODO API")
	})

	http.HandleFunc("/register",handlers.RegisterHandler)
	http.HandleFunc("/login",handlers.LoginHandler)
	
	fmt.Println("Server started at : 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
