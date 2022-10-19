package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/freddyouellette/duolingo-text-extractor/pkg/service/duolingo_extractor"
	"github.com/freddyouellette/duolingo-text-extractor/pkg/service/language_detector"
	"github.com/freddyouellette/duolingo-text-extractor/pkg/service/text_extractor"
	"github.com/stretchr/testify/assert"
)

type IntegrationTest struct {
	FilePath       string
	OriginalText   string
	TranslatedText string
}

func TestIntegration(t *testing.T) {
	tests := []IntegrationTest{
		{
			FilePath:       "./data/i2.png",
			OriginalText:   "Quest'automobile è come nuova.",
			TranslatedText: "This automobile is like new.",
		},
		{
			FilePath:       "./data/i4.jpg",
			OriginalText:   "La mia amica mi ha fatto conoscere mio marito.",
			TranslatedText: "My friend introduced me to my husband.",
		},
		{
			FilePath:       "./data/i5.jpg",
			OriginalText:   "Mio nonno è di età avanzata.",
			TranslatedText: "My grandfather is of an advanced age.",
		},
		{
			FilePath:       "./data/i6.jpg",
			OriginalText:   "Non so più a chi credere.",
			TranslatedText: "I do not know who to believe anymore.",
		},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Integration Test #%d (%s)", i, test.FilePath)
		t.Run(testName, func(t *testing.T) {
			textExtractor := new(text_extractor.AwsRekognition)
			languageDetector := new(language_detector.AwsComprehend)

			fileBytes, err := os.ReadFile(test.FilePath)
			assert.NoError(t, err)

			originalText, translatedText, err := duolingo_extractor.ExtractTranslations(textExtractor, languageDetector, fileBytes)
			assert.NoError(t, err)
			assert.Equal(t, test.OriginalText, originalText)
			assert.Equal(t, test.TranslatedText, translatedText)
		})
	}
}
