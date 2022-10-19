package duolingo_extractor

import (
	"strings"

	"github.com/freddyouellette/duolingo-text-extractor/pkg/service/language_detector"
	"github.com/freddyouellette/duolingo-text-extractor/pkg/service/text_extractor"
)

func ExtractTranslations(
	textExtractorService text_extractor.TextExtractorService,
	languageDetectionService language_detector.LanguageDetector,
	inputBytes []byte,
) (originalString string, translatedString string, err error) {
	lines, err := textExtractorService.ExtractText(inputBytes)
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
			originalLanguage, err = languageDetectionService.DetectLanguage([]byte(line))
			if err != nil {
				return "", "", err
			}
			originalLines = append(originalLines, line)
		} else {
			thisLanguage, err := languageDetectionService.DetectLanguage([]byte(line))
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
