package output_formatting

import (
	"fmt"
	"testing"

	"github.com/freddyouellette/duolingo-text-extractor/pkg/data"
	"github.com/stretchr/testify/assert"
)

type TextOutputterTest struct {
	Name   string
	Input  data.Translation
	Output string
}

func TestTextOutputter(t *testing.T) {
	tests := []TextOutputterTest{
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
			Output: "Quest'automobile è come nuova.\n",
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
			Output: "This car is like new.\nQuest'automobile è come nuova.\n",
		},
		{
			Name: "Nil Texts",
			Input: data.Translation{
				Texts: nil,
			},
			Output: "",
		},
		{
			Name: "Zero Texts",
			Input: data.Translation{
				Texts: []data.Text{},
			},
			Output: "",
		},
	}

	TextOutputter := new(TextOutputter)

	for i, test := range tests {
		t.Run(fmt.Sprintf("Text Outputter Test %d: %s", i, test.Name), func(t *testing.T) {
			thisOutput, err := TextOutputter.Render(test.Input)
			assert.NoError(t, err)
			assert.Equal(t, test.Output, thisOutput)
		})
	}
}
