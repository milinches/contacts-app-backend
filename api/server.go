package api

import (
	"log"
	"net/http"
	"os"

	"github.com/milinches/contacts-app-backend/api/router"
)

func Run() {
	// source .env
	port := os.Getenv("PORT")
	log.Println("Starting server...")

	r := router.NEW()

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
