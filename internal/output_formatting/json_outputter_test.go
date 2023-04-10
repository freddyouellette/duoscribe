package output_formatting

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/freddyouellette/duolingo-text-extractor/internal/models"
	"github.com/stretchr/testify/assert"
)

type JsonOutputterTest struct {
	name           string
	encodeJsonFunc func(v any) ([]byte, error)
	input          []models.Text
	expectedOutput string
	expectedError  error
}

func TestJsonOutputter(t *testing.T) {
	tests := []JsonOutputterTest{
		{
			name:           "One Text",
			encodeJsonFunc: json.Marshal,
			input: []models.Text{
				{
					Language: "it",
					Text:     "Quest'automobile è come nuova.",
				},
			},
			expectedOutput: "[{\"Language\":\"it\",\"Text\":\"Quest'automobile è come nuova.\"}]",
			expectedError:  nil,
		},
		{
			name:           "Two Texts",
			encodeJsonFunc: json.Marshal,
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
			expectedOutput: "[{\"Language\":\"en\",\"Text\":\"This car is like new.\"},{\"Language\":\"it\",\"Text\":\"Quest'automobile è come nuova.\"}]",
			expectedError:  nil,
		},
		{
			name: "Json Error",
			encodeJsonFunc: func(v any) ([]byte, error) {
				return nil, errors.New("json marshal error")
			},
			input: []models.Text{
				{
					Language: "it",
					Text:     "Quest'automobile è come nuova.",
				},
			},
			expectedOutput: "",
			expectedError:  errEncodingJson,
		},
		{
			name:           "Nil Texts",
			encodeJsonFunc: json.Marshal,
			input:          nil,
			expectedOutput: "",
			expectedError:  nil,
		},
		{
			name:           "Zero Texts",
			encodeJsonFunc: json.Marshal,
			input:          []models.Text{},
			expectedOutput: "",
			expectedError:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			jsonOutputter := NewJsonOutputter(test.encodeJsonFunc)
			actualOutput, actualErr := jsonOutputter.Render(test.input)
			assert.ErrorIs(t, actualErr, test.expectedError)
			assert.Equal(t, test.expectedOutput, actualOutput)
		})
	}
}
