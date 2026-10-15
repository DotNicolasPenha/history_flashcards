package main

import (
	"log"
	"net/http"

	"github.com/dotnicolaspenha/historyflashcards/internal/flashcard"
)

func main() {
	mux := http.NewServeMux()

	// mux.HandleFunc("GET /v1/flashcards", flashcard.GetFlashcardsHandler)
	// mux.HandleFunc("GET /v1/flashcard/{id}", flashcard.GetOneFlashcardByIdHandler)
	mux.HandleFunc("GET /v1/random_flashcard", flashcard.GetRandomFlashcardHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
