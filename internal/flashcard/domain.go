package flashcard

type Flashcard struct {
	ID               string `json:"ID"`
	Title            string `json:"title"`
	Question         string `json:"question"`
	QuestionResponse string `json:"question_response"`
}

type FlashcardStore interface{
	GetAll() ([]Flashcard, error)
	GetRandom() (Flashcard, error)
	GetOneByID(id string) (Flashcard,error)
}
