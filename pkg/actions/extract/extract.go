package extract

import "github.com/freddyouellette/duolingo-text-extractor/pkg/models"

type TextExtractor interface {
	ExtractText(inputBytes []byte) ([]string, error)
}

type LanguageDetector interface {
	DetectLanguage(inputBytes []byte) (string, error)
}

type TextCleaner interface {
	CleanText(text string) (string, error)
}

type TextCondenser interface {
	Condense(texts []models.Text) ([]models.Text, error)
}

type Outputter interface {
	Render(output []models.Text) (string, error)
}

type Action struct {
	TextExtractor    TextExtractor
	LanguageDetector LanguageDetector
	TextCleaner      TextCleaner
	TextCondenser    TextCondenser
	Outputter        Outputter
}

func (a *Action) Extract(imageBytes []byte) error {
	lines, err := a.TextExtractor.ExtractText(imageBytes)
	if err != nil {
		return err
	}

	var texts []models.Text
	for _, line := range lines {
		line, err := a.TextCleaner.CleanText(line)
		if err != nil {
			return err
		}
		if line == "" {
			continue
		}

		language, err := a.LanguageDetector.DetectLanguage([]byte(line))
		if err != nil {
			return err
		}

		texts = append(texts, models.Text{
			Language: language,
			Text:     line,
		})
	}

	texts, err = a.TextCondenser.Condense(texts)
	if err != nil {
		return err
	}

	a.Outputter.Render(texts)

	return nil
}
