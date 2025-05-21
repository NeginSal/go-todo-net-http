package handlers

import (
	"fmt"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "just POST request is allowed!", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, "sign up is successful")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "just POST request is allowed!", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "login is successful")
}
