package flashcard

type Flashcard struct {
	ID               string `json:"ID"`
	Title            string `json:"title"`
	Question         string `json:"question"`
	QuestionResponse string `json:"question_response"`
}

type FlashCardStore interface {
	GetRandom() (*Flashcard, error)
}
