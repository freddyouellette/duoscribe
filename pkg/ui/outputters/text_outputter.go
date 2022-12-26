package output_formatting

import (
	"fmt"

	"github.com/freddyouellette/duolingo-text-extractor/pkg/data"
)

type TextOutputter struct{}

func (o *TextOutputter) Render(output data.Translation) (string, error) {
	if output.Texts == nil {
		output.Texts = make([]data.Text, 0)
	}

	outputString := ""
	for _, text := range output.Texts {
		outputString += fmt.Sprintln(text.Text)
	}

	return outputString, nil
}
