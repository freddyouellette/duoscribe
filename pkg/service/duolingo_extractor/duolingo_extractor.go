package duolingo_extractor

import (
	"errors"
	"strings"

	"github.com/freddyouellette/duolingo-text-extractor/pkg/service/text_extractor"
)

func ExtractTranslations(service text_extractor.TextExtractorService, inputBytes []byte) ([]string, error) {
	lines, err := service.ExtractText(inputBytes)
	if err != nil {
		return []string{}, err
	}

	var newLines []string
	for _, line := range lines {
		if strings.ToLower(line) != "duolingo" {
			newLines = append(newLines, line)
		}
	}

	if len(newLines) != 2 {
		return []string{}, errors.New("provided image does not have 2 lines of text")
	}

	return newLines, err
}
