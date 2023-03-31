package output_formatting

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/freddyouellette/duolingo-text-extractor/internal/models"
)

var errMarshalling = errors.New("error marshalling the Text object")

// JsonOutputter outputs an array of Texts as JSON.
type JsonOutputter struct{}

// Render will return the array of Texts in JSON format.
func (o *JsonOutputter) Render(output []models.Text) (string, error) {
	if len(output) == 0 {
		return "", nil
	}

	outputBytes, err := json.Marshal(output)
	if err != nil {
		return "", fmt.Errorf("%w: %s", errMarshalling, err)
	}
	return string(outputBytes), nil
}
