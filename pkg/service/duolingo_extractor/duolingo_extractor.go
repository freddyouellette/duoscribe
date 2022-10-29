package duolingo_extractor

import (
	"fmt"
	"strings"

	"github.com/freddyouellette/duolingo-text-extractor/pkg/service/language_detector"
	"github.com/freddyouellette/duolingo-text-extractor/pkg/service/text_extractor"
)

type TextOrganizerService struct {
	textExtractorService     text_extractor.TextExtractorService
	languageDetectionService language_detector.LanguageDetector
}

type TextOrganizerOutput struct {
	Texts []Text
}

type Text struct {
	Language string
	Text     string
}

func NewTextOrganizerService(
	textExtractorService text_extractor.TextExtractorService,
	languageDetectionService language_detector.LanguageDetector,
) TextOrganizerService {
	return TextOrganizerService{
		textExtractorService:     textExtractorService,
		languageDetectionService: languageDetectionService,
	}
}

func (to *TextOrganizerService) ExtractTranslations(inputBytes []byte) (originalString string, translatedString string, err error) {
	lines, err := to.textExtractorService.ExtractText(inputBytes)
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
			originalLanguage, err = to.languageDetectionService.DetectLanguage([]byte(line))
			if err != nil {
				return "", "", err
			}
			originalLines = append(originalLines, line)
		} else {
			thisLanguage, err := to.languageDetectionService.DetectLanguage([]byte(line))
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
