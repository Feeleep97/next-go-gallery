package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Feeleep97/next-go-gallery/models"
)

var users = []models.User{
	{
		ID: 1, Name: "Alice", Email: "alice@example.com",
		ID: 2, Name: "Bob", Email: "bob@example.com",
	},
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var newUser models.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	newUser.id = len(users) + 1
	users = append(users, newUser)
	json.NewEncoder(w).Encode(newUser)
}
