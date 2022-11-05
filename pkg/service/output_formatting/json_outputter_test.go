package output_formatting

import (
	"fmt"
	"testing"

	"github.com/freddyouellette/duolingo-text-extractor/pkg/service/text_cleaning"
	"github.com/stretchr/testify/assert"
)

type JsonOutputterTest struct {
	Input  text_cleaning.TextCleanerOutput
	Output string
}

func TestJsonOutputter(t *testing.T) {
	tests := []JsonOutputterTest{
		{
			Input: text_cleaning.TextCleanerOutput{
				Texts: []text_cleaning.Text{
					{
						Language: "it",
						Text:     "Quest'automobile è come nuova.",
					},
				},
			},
			Output: "{\"Texts\":[{\"Language\":\"it\",\"Text\":\"Quest'automobile è come nuova.\"}]}",
		},
	}

	jsonOutputter := new(JsonOutputter)

	for i, test := range tests {
		t.Run(fmt.Sprintf("JSON Outputter Test %d", i), func(t *testing.T) {
			thisOutput, err := jsonOutputter.Render(test.Input)
			assert.NoError(t, err)
			assert.Equal(t, test.Output, thisOutput)
		})
	}
}
