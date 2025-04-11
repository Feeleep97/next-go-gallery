package routes

import (
	"github.com/Feeleep97/next-go-gallery/handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("api/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("api/users", handlers.CreateUser).Methods("POST")
}
