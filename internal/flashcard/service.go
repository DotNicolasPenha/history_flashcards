package flashcard

type FlashCardService struct {
	store FlashCardStore
}

func NewFlashCardService(store FlashCardStore) *FlashCardService {
	return &FlashCardService{
		store: store,
	}
}

func (s *FlashCardService) GetRandom() (*Flashcard, error) {
	return s.store.GetRandom()
}
