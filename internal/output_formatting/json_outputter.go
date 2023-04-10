package output_formatting

import (
	"errors"
	"fmt"

	"github.com/freddyouellette/duolingo-text-extractor/internal/models"
)

// JsonOutputter outputs an array of Texts as JSON.
type JsonOutputter struct {
	encodeJson func(v any) ([]byte, error)
}

func NewJsonOutputter(encodeJsonFunc func(v any) ([]byte, error)) *JsonOutputter {
	return &JsonOutputter{
		encodeJson: encodeJsonFunc,
	}
}

var errEncodingJson = errors.New("error encoding the Text object to json")

// Render will return the array of Texts in JSON format.
func (o *JsonOutputter) Render(output []models.Text) (string, error) {
	if len(output) == 0 {
		return "", nil
	}

	outputBytes, err := o.encodeJson(output)
	if err != nil {
		return "", fmt.Errorf("%w: %s", errEncodingJson, err)
	}
	return string(outputBytes), nil
}
