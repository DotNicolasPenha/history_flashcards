package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"

	"github.com/dotnicolaspenha/historyflashcards/internal/flashcard"
	"github.com/dotnicolaspenha/historyflashcards/internal/util"
	_ "github.com/dotnicolaspenha/historyflashcards/internal/util"
)

func main() {
	port := ":8080"
	mux := http.NewServeMux()

	// mux.HandleFunc("GET /v1/flashcards", flashcard.GetFlashcardsHandler)
	// mux.HandleFunc("GET /v1/flashcard/{id}", flashcard.GetOneFlashcardByIdHandler)
	mux.HandleFunc("GET /v1/random_flashcard", flashcard.GetRandomFlashcardHandler)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Warn("Path not found.",
			slog.String("path", r.URL.Path),
			slog.String("remote_addr", r.RemoteAddr))

		util.WriteJson(w, 404, util.Response{
			Ok:    false,
			Error: errors.New("Path not found.").Error(),
		})
	})

	slog.Info("Server initialized.", slog.String("PORT", port))

	if err := http.ListenAndServe(port, mux); err != nil {
		slog.Error("Error to init server.", slog.Any("error", err))
		os.Exit(0)
	}
}
