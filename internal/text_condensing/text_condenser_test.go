package text_condensing

import (
	"testing"

	"github.com/freddyouellette/duolingo-text-extractor/internal/models"
	"github.com/stretchr/testify/assert"
)

type textCondenserTest struct {
	Name           string
	Input          []models.Text
	ExpectedOutput []models.Text
	ExpectedError  error
}

func TestRekognitionExtractFromFile(t *testing.T) {
	tests := []textCondenserTest{
		{
			Name: "No condensing needed",
			Input: []models.Text{
				{
					Text:     "text",
					Language: "en",
				},
			},
			ExpectedOutput: []models.Text{
				{
					Text:     "text",
					Language: "en",
				},
			},
			ExpectedError: nil,
		},
		{
			Name: "Two into one",
			Input: []models.Text{
				{
					Text:     "text1",
					Language: "en",
				},
				{
					Text:     "text2",
					Language: "en",
				},
			},
			ExpectedOutput: []models.Text{
				{
					Text:     "text1 text2",
					Language: "en",
				},
			},
			ExpectedError: nil,
		},
		{
			Name: "Three into two",
			Input: []models.Text{
				{
					Text:     "text1",
					Language: "en",
				},
				{
					Text:     "text2",
					Language: "it",
				},
				{
					Text:     "text3",
					Language: "it",
				},
			},
			ExpectedOutput: []models.Text{
				{
					Text:     "text1",
					Language: "en",
				},
				{
					Text:     "text2 text3",
					Language: "it",
				},
			},
			ExpectedError: nil,
		},
		{
			Name: "Four into two",
			Input: []models.Text{
				{
					Text:     "text1",
					Language: "en",
				},
				{
					Text:     "text2",
					Language: "it",
				},
				{
					Text:     "text3",
					Language: "en",
				},
				{
					Text:     "text4",
					Language: "it",
				},
			},
			ExpectedOutput: []models.Text{
				{
					Text:     "text1 text3",
					Language: "en",
				},
				{
					Text:     "text2 text4",
					Language: "it",
				},
			},
			ExpectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			textCondenser := new(TextCondenser)
			actualOutput, err := textCondenser.Condense(test.Input)

			if test.ExpectedError != nil {
				assert.ErrorAs(t, err, test.ExpectedError)
			} else {
				assert.Equal(t, test.ExpectedOutput, actualOutput)
			}
		})
	}
}
