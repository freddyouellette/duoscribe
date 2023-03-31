package text_condensing

import "github.com/freddyouellette/duolingo-text-extractor/internal/models"

// TextCondenser condenses texts so that each language only has one string.
type TextCondenser struct{}

func NewTextCondenser() *TextCondenser {
	return &TextCondenser{}
}

// Condense takes an array of Texts and condenses them, so that each language has only one string.
// It will join each string with spaces.
func (t *TextCondenser) Condense(texts []models.Text) ([]models.Text, error) {
	languageTexts := make(map[string]string)
	for _, text := range texts {
		if _, ok := languageTexts[text.Language]; !ok {
			languageTexts[text.Language] = ""
		}
		if len(languageTexts[text.Language]) > 0 {
			languageTexts[text.Language] += " "
		}
		languageTexts[text.Language] += text.Text
	}

	var returnTexts []models.Text
	for language, text := range languageTexts {
		returnTexts = append(returnTexts, models.Text{
			Language: language,
			Text:     text,
		})
	}

	return returnTexts, nil
}
