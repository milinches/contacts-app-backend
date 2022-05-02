package api

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/milinches/contacts-app-backend/api/router"
)

func Run() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error loading .env file...")
	}

	port := os.Getenv("PORT")
	log.Println("Starting server...")

	r := router.NEW()

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
