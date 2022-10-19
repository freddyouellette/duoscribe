package duolingo_extractor

import (
	"testing"

	"github.com/freddyouellette/duolingo-text-extractor/pkg/service/language_detector"
	"github.com/freddyouellette/duolingo-text-extractor/pkg/service/text_extractor"
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

			textExtractorServiceMock := new(text_extractor.TextExtractorServiceMock)
			textExtractorServiceMock.On("ExtractText", bytesGiven).Return(test.ExtractedLines, nil)

			languageDetectorServiceMock := new(language_detector.LanguageDetectorMock)
			for i, line := range test.ExtractedLines {
				languageDetectorServiceMock.On("DetectLanguage", []byte(line)).Return(test.ExtractedLanguages[i], nil)
			}

			originalString, translatedString, err := ExtractTranslations(textExtractorServiceMock, languageDetectorServiceMock, bytesGiven)
			assert.ErrorIs(t, err, test.expectedError)

			assert.NoError(t, err)
			assert.Equal(t, test.ExpectedOriginal, originalString)
			assert.Equal(t, test.ExpectedTranslated, translatedString)
		})
	}
}
