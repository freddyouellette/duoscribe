package output_formatting

import (
	"fmt"

	"github.com/freddyouellette/duolingo-text-extractor/pkg/models"
)

type TextOutputter struct{}

func (o *TextOutputter) Render(output []models.Text) (string, error) {
	outputString := ""
	for _, text := range output {
		outputString += fmt.Sprintln(text.Text)
	}

	return outputString, nil
}
