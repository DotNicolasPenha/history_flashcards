package flashcard

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/dotnicolaspenha/historyflashcards/internal/util"
)

type FlashCardHandler struct {
	service FlashCardService
}

func NewFlashCardHandler(service *FlashCardService) *FlashCardHandler {
	return &FlashCardHandler{
		service: *service,
	}
}

func (f *FlashCardHandler) GetRandomHandler(w http.ResponseWriter, r *http.Request) {
	flashcard, err := f.service.GetRandom()

	if err != nil {
		slog.Error("Error to get random flashcard.", slog.Any("Error", err))
		util.WriteJson(w, 500, util.Response{
			Ok:    false,
			Error: errors.New("Error to get random flashcard.").Error(),
		})
	}

	slog.Info("GET v1/random_flashcard",
		slog.String("flashcard_id", flashcard.ID),
		slog.String("remote_addr", r.RemoteAddr))

	util.WriteJson(w, 200, util.Response{
		Ok:   true,
		Data: flashcard,
	})
}
