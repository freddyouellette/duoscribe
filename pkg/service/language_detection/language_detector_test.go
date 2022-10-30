package language_detection

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type LanguageDetectorTest struct {
	Name             string
	InputString      string
	ExpectedLanguage string
}

func TestLanguageDetector(t *testing.T) {
	tests := []LanguageDetectorTest{
		{
			Name:             "Italian 1",
			InputString:      "Lo può ripetere per favore?",
			ExpectedLanguage: "it",
		},
		{
			Name:             "English 1",
			InputString:      "Can you repeat it please?",
			ExpectedLanguage: "en",
		},
		{
			Name:             "Non-Language 1",
			InputString:      "duolingo",
			ExpectedLanguage: "en",
		},
		{
			Name:             "Ambiguous 1",
			InputString:      "pasta",
			ExpectedLanguage: "en",
		},
	}

	languageDetectorConcretions := map[string]LanguageDetector{
		"AWS Comprehend": new(AwsComprehend),
	}
	for concretionName, concretion := range languageDetectorConcretions {
		for _, test := range tests {
			testName := fmt.Sprintf("%s - %s", concretionName, test.Name)
			t.Run(testName, func(t *testing.T) {
				detectedLanguage, err := concretion.DetectLanguage([]byte(test.InputString))
				assert.NoError(t, err)
				assert.Equal(t, test.ExpectedLanguage, detectedLanguage)
			})
		}
	}
}
