package flashcard

import "math/rand/v2"

type MemoryFlashCardStore struct {
	flashcards []Flashcard
}

func NewMemoryFlashCardStore() *MemoryFlashCardStore {
	return &MemoryFlashCardStore{
		flashcards: []Flashcard{
			{
				ID:               "1",
				Title:            "Causas da Revolução Francesa",
				Question:         "Quais foram as principais causas da Revolução Francesa?",
				QuestionResponse: "A desigualdade social, a crise financeira, os altos impostos sobre o Terceiro Estado e as ideias iluministas.",
			},
			{
				ID:               "2",
				Title:            "Início da Revolução",
				Question:         "Em que ano começou a Revolução Francesa?",
				QuestionResponse: "Em 1789, durante uma grave crise política, econômica e social na França.",
			},
			{
				ID:               "3",
				Title:            "Queda da Bastilha",
				Question:         "O que aconteceu em 14 de julho de 1789?",
				QuestionResponse: "A população de Paris tomou a Bastilha, prisão que simbolizava o poder absolutista do rei.",
			},
			{
				ID:               "4",
				Title:            "Sociedade Francesa",
				Question:         "Quais eram os três estados da sociedade francesa?",
				QuestionResponse: "Primeiro Estado: clero; Segundo Estado: nobreza; Terceiro Estado: burguesia, trabalhadores urbanos e camponeses.",
			},
			{
				ID:               "5",
				Title:            "Antigo Regime",
				Question:         "O que era o Antigo Regime na França?",
				QuestionResponse: "Era o sistema político e social baseado na monarquia absolutista e nos privilégios do clero e da nobreza.",
			},
			{
				ID:               "6",
				Title:            "Iluminismo",
				Question:         "Como o Iluminismo influenciou a Revolução Francesa?",
				QuestionResponse: "Defendeu ideias como liberdade, igualdade perante a lei, soberania popular e crítica ao absolutismo.",
			},
			{
				ID:               "7",
				Title:            "Declaração dos Direitos",
				Question:         "O que foi a Declaração dos Direitos do Homem e do Cidadão?",
				QuestionResponse: "Um documento de 1789 que proclamou direitos como liberdade e igualdade perante a lei, embora não fossem aplicados igualmente a todos.",
			},
			{
				ID:               "8",
				Title:            "Fim da Monarquia",
				Question:         "O que aconteceu com a monarquia francesa em 1792?",
				QuestionResponse: "A monarquia foi abolida e a França tornou-se uma república.",
			},
			{
				ID:               "9",
				Title:            "Período do Terror",
				Question:         "O que foi o Período do Terror?",
				QuestionResponse: "Uma fase de 1793 a 1794 marcada por repressão política e execuções em massa, durante a qual os jacobinos tiveram papel central.",
			},
			{
				ID:               "10",
				Title:            "Fim da Revolução",
				Question:         "Qual acontecimento é geralmente considerado o fim da Revolução Francesa?",
				QuestionResponse: "O golpe de 18 de Brumário, em novembro de 1799, que levou Napoleão Bonaparte ao poder.",
			},
		},
	}
}

func (store *MemoryFlashCardStore) GetRandom() (*Flashcard, error) {
	return &store.flashcards[rand.IntN(len(store.flashcards))], nil
}
