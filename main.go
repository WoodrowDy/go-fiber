package main

import (
	"github.com/joho/godotenv"
	"go-fiber/models"
	"go-fiber/routes"
	"log"
	"os"
)

func main() {
	// .env 로드
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// port 가져오기
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	models.InitDB()
	app := routes.Routes()

	log.Fatal(app.Listen(`0.0.0.0:` + port))
}
