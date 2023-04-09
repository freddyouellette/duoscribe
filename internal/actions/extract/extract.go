package extract

import (
	"errors"
	"fmt"

	"github.com/freddyouellette/duolingo-text-extractor/internal/models"
)

// TextExtractor extracts text from an image
type TextExtractor interface {
	ExtractText(inputBytes []byte) ([]string, error)
}

// LanguageDetector detects the language of a string
type LanguageDetector interface {
	DetectLanguage(inputBytes []byte) (string, error)
}

// TextCleaner removes unwanted text from a string
type TextCleaner interface {
	CleanText(text string) (string, error)
}

// TextCondenser condenses Text to one string per language
type TextCondenser interface {
	Condense(texts []models.Text) ([]models.Text, error)
}

// Outputter formats the output of an array of Texts
type Outputter interface {
	Render(output []models.Text) (string, error)
}

// Action will extract text from image bytes and outputs the result to the console.
type Action struct {
	TextExtractor    TextExtractor
	TextCleaner      TextCleaner
	LanguageDetector LanguageDetector
	TextCondenser    TextCondenser
	Outputter        Outputter
}

var (
	errTextExtractor    = errors.New("extracting text from image")
	errTextCleaner      = errors.New("cleaning text")
	errLanguageDetector = errors.New("detecting language")
	errTextCondenser    = errors.New("condensing text")
	errOutputter        = errors.New("outputting text")
)

// Extract transcribes text from an image and outputs the result to the console.
func (a *Action) Extract(imageBytes []byte) (string, error) {
	lines, err := a.TextExtractor.ExtractText(imageBytes)
	if err != nil {
		return "", fmt.Errorf("%w: %s", errTextExtractor, err)
	}

	var texts []models.Text
	for _, line := range lines {
		line, err := a.TextCleaner.CleanText(line)
		if err != nil {
			return "", fmt.Errorf("%w: %s", errTextCleaner, err)
		}
		if line == "" {
			continue
		}

		language, err := a.LanguageDetector.DetectLanguage([]byte(line))
		if err != nil {
			return "", fmt.Errorf("%w: %s", errLanguageDetector, err)
		}

		texts = append(texts, models.Text{
			Language: language,
			Text:     line,
		})
	}

	texts, err = a.TextCondenser.Condense(texts)
	if err != nil {
		return "", fmt.Errorf("%w: %s", errTextCondenser, err)
	}

	out, err := a.Outputter.Render(texts)
	if err != nil {
		return "", fmt.Errorf("%w: %s", errOutputter, err)
	}

	return out, nil
}
