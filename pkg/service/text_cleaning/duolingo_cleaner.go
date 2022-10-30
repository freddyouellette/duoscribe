package text_cleaning

import (
	"fmt"
	"strings"

	"github.com/freddyouellette/duolingo-text-extractor/pkg/service/language_detection"
	"github.com/freddyouellette/duolingo-text-extractor/pkg/service/text_extraction"
)

type DuolingoCleaner struct {
	TextExtractor    text_extraction.TextExtractor
	languageDetector language_detection.LanguageDetector
}

func NewDuolingoCleaner(
	TextExtractor text_extraction.TextExtractor,
	languageDetector language_detection.LanguageDetector,
) TextCleaner {
	return &DuolingoCleaner{
		TextExtractor:    TextExtractor,
		languageDetector: languageDetector,
	}
}

func (to *DuolingoCleaner) ExtractTranslations(inputBytes []byte) (originalString string, translatedString string, err error) {
	lines, err := to.TextExtractor.ExtractText(inputBytes)
	if err != nil {
		return "", "", err
	}

	var originalLanguage string
	var originalLines []string
	var translatedLines []string
	for _, line := range lines {
		if strings.ToLower(line) == "duolingo" {
			// strip out the duolingo logo
			continue
		}

		if originalLanguage == "" {
			originalLanguage, err = to.languageDetector.DetectLanguage([]byte(line))
			if err != nil {
				return "", "", err
			}
			originalLines = append(originalLines, line)
		} else {
			thisLanguage, err := to.languageDetector.DetectLanguage([]byte(line))
			if err != nil {
				return "", "", err
			}
			if thisLanguage == originalLanguage {
				originalLines = append(originalLines, line)
			} else {
				translatedLines = append(translatedLines, line)
			}
		}
	}

	return strings.Join(originalLines, " "), strings.Join(translatedLines, " "), err
}

func (too *TextOrganizerOutput) GetLanguageText(language string) (string, error) {
	for _, text := range too.Texts {
		if text.Language == language {
			return text.Text, nil
		}
	}
	return "", fmt.Errorf("language \"%s\" not found", language)
}
