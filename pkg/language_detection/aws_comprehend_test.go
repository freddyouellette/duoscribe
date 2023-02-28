package language_detection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type LanguageDetectorTest struct {
	name             string
	inputString      string
	expectedLanguage string
}

func TestLanguageDetector(t *testing.T) {
	tests := []LanguageDetectorTest{
		{
			name:             "Italian 1",
			inputString:      "Lo può ripetere per favore?",
			expectedLanguage: "it",
		},
		{
			name:             "English 1",
			inputString:      "Can you repeat it please?",
			expectedLanguage: "en",
		},
		{
			name:             "Non-Language 1",
			inputString:      "duolingo",
			expectedLanguage: "en",
		},
		{
			name:             "Ambiguous 1",
			inputString:      "pasta",
			expectedLanguage: "en",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			languageDetector := &AwsComprehend{}
			detectedLanguage, err := languageDetector.DetectLanguage([]byte(test.inputString))
			assert.NoError(t, err)
			assert.Equal(t, test.expectedLanguage, detectedLanguage)
		})
	}
}
