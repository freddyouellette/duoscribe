package text_condensing

import (
	"sort"

	"github.com/freddyouellette/duolingo-text-extractor/internal/models"
)

// TextCondenser condenses texts so that each language only has one string.
type TextCondenser struct{}

func NewTextCondenser() *TextCondenser {
	return &TextCondenser{}
}

// Condense takes an array of Texts and condenses them, so that each language has only one string.
// It will join each string with spaces.
func (t *TextCondenser) Condense(texts []models.Text) ([]models.Text, error) {
	languages := make([]string, 0)
	languageTexts := make(map[string]string)
	for _, text := range texts {
		if _, ok := languageTexts[text.Language]; !ok {
			languageTexts[text.Language] = ""
			languages = append(languages, text.Language)
		}
		if len(languageTexts[text.Language]) > 0 {
			languageTexts[text.Language] += " "
		}
		languageTexts[text.Language] += text.Text
	}

	sort.Strings(languages)
	var returnTexts []models.Text
	for _, language := range languages {
		returnTexts = append(returnTexts, models.Text{
			Language: language,
			Text:     languageTexts[language],
		})
	}

	return returnTexts, nil
}
