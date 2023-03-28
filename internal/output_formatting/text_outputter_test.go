package output_formatting

import (
	"fmt"
	"testing"

	"github.com/freddyouellette/duolingo-text-extractor/internal/models"
	"github.com/stretchr/testify/assert"
)

type TextOutputterTest struct {
	name   string
	input  []models.Text
	output string
}

func TestTextOutputter(t *testing.T) {
	tests := []TextOutputterTest{
		{
			name: "One Text",
			input: []models.Text{
				{
					Language: "it",
					Text:     "Quest'automobile è come nuova.",
				},
			},
			output: "Quest'automobile è come nuova.\n",
		},
		{
			name: "Two Texts",
			input: []models.Text{
				{
					Language: "en",
					Text:     "This car is like new.",
				},
				{
					Language: "it",
					Text:     "Quest'automobile è come nuova.",
				},
			},
			output: "This car is like new.\nQuest'automobile è come nuova.\n",
		},
		{
			name:   "Nil Texts",
			input:  nil,
			output: "",
		},
		{
			name:   "Zero Texts",
			input:  []models.Text{},
			output: "",
		},
	}

	TextOutputter := new(TextOutputter)

	for i, test := range tests {
		t.Run(fmt.Sprintf("Text Outputter Test %d: %s", i, test.name), func(t *testing.T) {
			thisOutput, err := TextOutputter.Render(test.input)
			assert.NoError(t, err)
			assert.Equal(t, test.output, thisOutput)
		})
	}
}
