package main

import (
	"context"
	"log"
	"main/handlers"
	"main/middleware"
	"main/server"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v\n", err)
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")
	SUDO_TOKEN := os.Getenv("SUDO_TOKEN")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTsecret:   JWT_SECRET,
		DatabaseURL: DATABASE_URL,
		SudoToken:   SUDO_TOKEN,
	})

	if err != nil {
		log.Fatalf("Error creating server %v\n", err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	r.Use(middleware.CheckAuthMiddleware(s))

	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/create-user-being-root-ocult-method-with-password", handlers.SingUpHandler(s)).Methods(http.MethodPut)
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/me", handlers.MeHandler(s)).Methods(http.MethodGet)
}
