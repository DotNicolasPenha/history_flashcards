package main

import (
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/dotnicolaspenha/historyflashcards/internal/flashcard"
	"github.com/dotnicolaspenha/historyflashcards/internal/util"
	_ "github.com/dotnicolaspenha/historyflashcards/internal/util"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = ":8080"
	}

	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	repository := flashcard.NewMemoryFlashCardStore() // memory
	service := flashcard.NewFlashCardService(repository)
	handlers := flashcard.NewFlashCardHandler(service)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/random_flashcard", handlers.GetRandomHandler)
	mux.HandleFunc("/", util.NotFoundHandler)
	mux.HandleFunc("/health", util.HealthHandler)

	slog.Info("Server initialized.", slog.String("PORT", port))

	if err := http.ListenAndServe(port, mux); err != nil {
		slog.Error("Error to init server.", slog.Any("error", err))
		os.Exit(0)
	}
}
