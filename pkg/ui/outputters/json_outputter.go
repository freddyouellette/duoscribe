package output_formatting

import (
	"encoding/json"

	"github.com/freddyouellette/duolingo-text-extractor/pkg/data"
)

type JsonOutputter struct{}

func (o *JsonOutputter) Render(output data.Translation) (string, error) {

	if output.Texts == nil {
		output.Texts = make([]data.Text, 0)
	}

	outputBytes, err := json.Marshal(output)
	if err != nil {
		return "", err
	}
	return string(outputBytes), nil
}
