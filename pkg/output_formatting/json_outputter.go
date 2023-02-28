package output_formatting

import (
	"encoding/json"

	"github.com/freddyouellette/duolingo-text-extractor/pkg/models"
)

type JsonOutputter struct{}

func (o *JsonOutputter) Render(output []models.Text) (string, error) {
	if len(output) == 0 {
		return "", nil
	}

	outputBytes, err := json.Marshal(output)
	if err != nil {
		return "", err
	}
	return string(outputBytes), nil
}
