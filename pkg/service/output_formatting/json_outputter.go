package output_formatting

import (
	"encoding/json"

	"github.com/freddyouellette/duolingo-text-extractor/pkg/service/text_cleaning"
)

type JsonOutputter struct{}

func (o *JsonOutputter) Render(output text_cleaning.TextCleanerOutput) (string, error) {

	outputBytes, err := json.Marshal(output)
	if err != nil {
		return "", err
	}
	return string(outputBytes), nil
}
