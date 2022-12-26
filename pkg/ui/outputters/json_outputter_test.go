package output_formatting

import (
	"fmt"
	"testing"

	"github.com/freddyouellette/duolingo-text-extractor/pkg/data"
	"github.com/stretchr/testify/assert"
)

type JsonOutputterTest struct {
	Name   string
	Input  data.Translation
	Output string
}

func TestJsonOutputter(t *testing.T) {
	tests := []JsonOutputterTest{
		{
			Name: "One Text",
			Input: data.Translation{
				Texts: []data.Text{
					{
						Language: "it",
						Text:     "Quest'automobile è come nuova.",
					},
				},
			},
			Output: "{\"Texts\":[{\"Language\":\"it\",\"Text\":\"Quest'automobile è come nuova.\"}]}",
		},
		{
			Name: "Two Texts",
			Input: data.Translation{
				Texts: []data.Text{
					{
						Language: "en",
						Text:     "This car is like new.",
					},
					{
						Language: "it",
						Text:     "Quest'automobile è come nuova.",
					},
				},
			},
			Output: "{\"Texts\":[{\"Language\":\"en\",\"Text\":\"This car is like new.\"},{\"Language\":\"it\",\"Text\":\"Quest'automobile è come nuova.\"}]}",
		},
		{
			Name: "Nil Texts",
			Input: data.Translation{
				Texts: nil,
			},
			Output: "{\"Texts\":[]}",
		},
		{
			Name: "Zero Texts",
			Input: data.Translation{
				Texts: []data.Text{},
			},
			Output: "{\"Texts\":[]}",
		},
	}

	jsonOutputter := new(JsonOutputter)

	for i, test := range tests {
		t.Run(fmt.Sprintf("JSON Outputter Test %d: %s", i, test.Name), func(t *testing.T) {
			thisOutput, err := jsonOutputter.Render(test.Input)
			assert.NoError(t, err)
			assert.Equal(t, test.Output, thisOutput)
		})
	}
}
