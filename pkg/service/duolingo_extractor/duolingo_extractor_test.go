package duolingo_extractor

import (
	"testing"

	"github.com/freddyouellette/duolingo-text-extractor/pkg/service/text_extractor"
	"github.com/stretchr/testify/assert"
)

type ExtractTranslationsTest struct {
	Name          string
	GivenLines    []string
	ExpectedLines []string
}

func TestExtractTranslations(t *testing.T) {
	var bytesGiven []byte

	tests := []ExtractTranslationsTest{
		{
			Name: "Exact match",
			GivenLines: []string{
				"Lo può ripetere per favore?",
				"Can you repeat it please?",
			},
			ExpectedLines: []string{
				"Lo può ripetere per favore?",
				"Can you repeat it please?",
			},
		},
		{
			Name: "With Duolingo",
			GivenLines: []string{
				"Lo può ripetere per favore?",
				"Can you repeat it please?",
				"duolingo",
			},
			ExpectedLines: []string{
				"Lo può ripetere per favore?",
				"Can you repeat it please?",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			serviceMock := new(text_extractor.TextExtractorServiceMock)
			serviceMock.On("ExtractText", bytesGiven).Return(test.GivenLines, nil)

			outputLines, err := ExtractTranslations(serviceMock, bytesGiven)

			assert.NoError(t, err)
			assert.Equal(t, test.ExpectedLines, outputLines)
		})
	}
}
