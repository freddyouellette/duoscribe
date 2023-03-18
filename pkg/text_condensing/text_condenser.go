package text_condensing

import "github.com/freddyouellette/duolingo-text-extractor/pkg/models"

type TextCondenser struct{}

func (t *TextCondenser) Condense(texts []models.Text) ([]models.Text, error) {
	var languageTexts map[string]string = make(map[string]string)
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
