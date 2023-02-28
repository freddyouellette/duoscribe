package output_formatting

import (
	"fmt"
	"testing"

	"github.com/freddyouellette/duolingo-text-extractor/pkg/models"
	"github.com/stretchr/testify/assert"
)

type JsonOutputterTest struct {
	name   string
	input  []models.Text
	output string
}

func TestJsonOutputter(t *testing.T) {
	tests := []JsonOutputterTest{
		{
			name: "One Text",
			input: []models.Text{
				{
					Language: "it",
					Text:     "Quest'automobile è come nuova.",
				},
			},
			output: "[{\"Language\":\"it\",\"Text\":\"Quest'automobile è come nuova.\"}]",
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
			output: "[{\"Language\":\"en\",\"Text\":\"This car is like new.\"},{\"Language\":\"it\",\"Text\":\"Quest'automobile è come nuova.\"}]",
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

	jsonOutputter := new(JsonOutputter)

	for i, test := range tests {
		t.Run(fmt.Sprintf("JSON Outputter Test %d: %s", i, test.name), func(t *testing.T) {
			thisOutput, err := jsonOutputter.Render(test.input)
			assert.NoError(t, err)
			assert.Equal(t, test.output, thisOutput)
		})
	}
}
