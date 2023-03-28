package extract

import (
	"fmt"

	"github.com/freddyouellette/duolingo-text-extractor/internal/models"
)

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
	TextCleaner      TextCleaner
	LanguageDetector LanguageDetector
	TextCondenser    TextCondenser
	Outputter        Outputter
}

const (
	errTextExtractor    = "extracting text from image"
	errTextCleaner      = "cleaning text"
	errLanguageDetector = "detecting language"
	errTextCondenser    = "condensing text"
	errOutputter        = "outputting text"
)

func (a *Action) Extract(imageBytes []byte) error {
	lines, err := a.TextExtractor.ExtractText(imageBytes)
	if err != nil {
		return fmt.Errorf("%s: %w", errTextExtractor, err)
	}

	var texts []models.Text
	for _, line := range lines {
		line, err := a.TextCleaner.CleanText(line)
		if err != nil {
			return fmt.Errorf("%s: %w", errTextCleaner, err)
		}
		if line == "" {
			continue
		}

		language, err := a.LanguageDetector.DetectLanguage([]byte(line))
		if err != nil {
			return fmt.Errorf("%s: %w", errLanguageDetector, err)
		}

		texts = append(texts, models.Text{
			Language: language,
			Text:     line,
		})
	}

	texts, err = a.TextCondenser.Condense(texts)
	if err != nil {
		return fmt.Errorf("%s: %w", errTextCondenser, err)
	}

	out, err := a.Outputter.Render(texts)
	if err != nil {
		return fmt.Errorf("%s: %w", errOutputter, err)
	}

	fmt.Print(out)

	return nil
}
