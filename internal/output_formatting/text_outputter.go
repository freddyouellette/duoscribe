package output_formatting

import (
	"fmt"

	"github.com/freddyouellette/duolingo-text-extractor/internal/models"
)

// TextOutputter can output an array of Texts in a simple text format.
type TextOutputter struct{}

// Render returns the Texts in a simple format with newlines, without the language.
func (o *TextOutputter) Render(output []models.Text) (string, error) {
	outputString := ""
	for _, text := range output {
		outputString += fmt.Sprintln(text.Text)
	}

	return outputString, nil
}
