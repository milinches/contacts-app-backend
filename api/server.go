package api

import (
	"log"
	"net/http"
	"os"

	"github.com/milinches/contacts-app-backend/api/router"
	"github.com/milinches/contacts-app-backend/api/auto"
)

// source .env
var port = os.Getenv("PORT")

// Starts a new server
func Run() {
	auto.Load()
	log.Println("Starting server...")

	r := router.NEW()

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
