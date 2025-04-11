package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/Feeleep97/next-go-gallery/backend"
)

func main() {
	r := mux.NewRouter()
	routers.registerRoutes(r)
	// println("this is backend-WIP")
	log.Println("🚀 Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
