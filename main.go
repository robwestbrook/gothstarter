package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/robwestbrook/gothstarter/handlers"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// Create router
	router := chi.NewMux()

	// Routes
	router.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	router.Get("/", handlers.Make(handlers.HandleHome))
	router.Get("/login", handlers.Make(handlers.HandleLogin))

	// Start server
	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listenAddr", listenAddr)
	http.ListenAndServe(listenAddr, router)
}
