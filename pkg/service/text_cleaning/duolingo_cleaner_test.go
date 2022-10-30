package text_cleaning

import (
	"testing"

	"github.com/freddyouellette/duolingo-text-extractor/pkg/service/language_detection"
	"github.com/freddyouellette/duolingo-text-extractor/pkg/service/text_extraction"
	"github.com/stretchr/testify/assert"
)

type ExtractTranslationsTest struct {
	Name               string
	ExtractedLines     []string
	ExtractedLanguages []string
	ExpectedOriginal   string
	ExpectedTranslated string
	expectedError      error
}

func TestExtractTranslations(t *testing.T) {
	var bytesGiven []byte

	tests := []ExtractTranslationsTest{
		{
			Name: "Exact match",
			ExtractedLines: []string{
				"Lo può ripetere per favore?",
				"Can you repeat it please?",
			},
			ExtractedLanguages: []string{
				"it",
				"en",
			},
			ExpectedOriginal:   "Lo può ripetere per favore?",
			ExpectedTranslated: "Can you repeat it please?",
			expectedError:      nil,
		},
		{
			Name: "With Duolingo",
			ExtractedLines: []string{
				"Lo può ripetere per favore?",
				"Can you repeat it please?",
				"duolingo",
			},
			ExtractedLanguages: []string{
				"it",
				"en",
				"en",
			},
			ExpectedOriginal:   "Lo può ripetere per favore?",
			ExpectedTranslated: "Can you repeat it please?",
			expectedError:      nil,
		},
		{
			Name:               "Empty",
			ExtractedLines:     []string{},
			ExtractedLanguages: []string{},
			ExpectedOriginal:   "",
			ExpectedTranslated: "",
			expectedError:      nil,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			TextExtractorMock := new(text_extraction.TextExtractorMock)
			TextExtractorMock.On("ExtractText", bytesGiven).Return(test.ExtractedLines, nil)

			languageDetectorMock := new(language_detection.LanguageDetectorMock)
			for i, line := range test.ExtractedLines {
				languageDetectorMock.On("DetectLanguage", []byte(line)).Return(test.ExtractedLanguages[i], nil)
			}

			textOrganizer := NewDuolingoCleaner(TextExtractorMock, languageDetectorMock)

			originalString, translatedString, err := textOrganizer.ExtractTranslations(bytesGiven)
			assert.ErrorIs(t, err, test.expectedError)

			assert.NoError(t, err)
			assert.Equal(t, test.ExpectedOriginal, originalString)
			assert.Equal(t, test.ExpectedTranslated, translatedString)
		})
	}
}
