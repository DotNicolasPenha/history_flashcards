package flashcard

import (
	"errors"
	"log/slog"
	"math/rand/v2"
	"net/http"

	"github.com/dotnicolaspenha/historyflashcards/internal/util"
)

func GetFlashcardsHandler(w http.ResponseWriter, r *http.Request) {
	// query := r.URL.Query()

	// page := 1
	// limit := 10
	// maxlimit := 100

	// if value := query.Get("page"); value != "" {
	// 	n, err := strconv.Atoi(value)

	// 	if err != nil || n < 1 {
	// 		util.WriteJson(w, 400, util.Response{
	// 			Ok:    false,
	// 			Error: errors.New("Page should be a positive int").Error(),
	// 		})

	// 		return
	// 	}
	// 	page = n
	// }

	// if value := query.Get("limit"); value != ""{
	// 	n,err:= strconv.Atoi(value)

	// 	if err != nil || n < 1 {
	// 		util.WriteJson(w, 400, util.Response{
	// 			Ok:    false,
	// 			Error: errors.New("Limit should be a positive int").Error(),
	// 		})

	// 		return
	// 	}

	// 	if n > maxlimit {
	// 		util.WriteJson(w, 400, util.Response{
	// 			Ok: false,
	// 			Error: errors.New("Max limit is ",maxlimit).Error(),
	// 		})

	// 		return
	// 	}

	// 	limit = n
	// }
	util.WriteJson(w, 200, util.Response{
		Ok:   true,
		Data: Flashcards,
	})
}

func GetRandomFlashcardHandler(w http.ResponseWriter, r *http.Request) {
	flashcard := Flashcards[rand.IntN(len(Flashcards))]

	slog.Info("GET v1/random_flashcard", slog.String("remote_addr", r.RemoteAddr))
	util.WriteJson(w, 200, util.Response{
		Ok:   true,
		Data: flashcard,
	})
}

func GetOneFlashcardByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	for _, f := range Flashcards {
		if f.ID == id {
			util.WriteJson(w, 200, util.Response{
				Ok:   true,
				Data: f,
			})
			return
		}
	}

	util.WriteJson(w, 404, util.Response{
		Ok:    false,
		Error: errors.New("Flashcard not found").Error(),
	})
}
