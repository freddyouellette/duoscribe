package text_extractor

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type RekognitionExtractFromFileTest struct {
	Name          string
	FilePath      string
	ExpectedLines []string
}

func TestRekognitionExtractFromFile(t *testing.T) {
	tests := []RekognitionExtractFromFileTest{
		{
			Name:     "1 line each",
			FilePath: "../../test/data/2_lines.jpg",
			ExpectedLines: []string{
				"Lo può ripetere per favore?",
				"Can you repeat it please?",
				"duolingo",
			},
		},
		{
			Name:     "3 lines",
			FilePath: "../../test/data/3_lines.jpg",
			ExpectedLines: []string{
				"Non so più a chi credere.",
				"I do not know who to",
				"believe anymore.",
				"duolingo",
			},
		},
		{
			Name:     "4 lines",
			FilePath: "../../test/data/4_lines.jpg",
			ExpectedLines: []string{
				"Due bicchieri di succo",
				"d'arancia, per piacere.",
				"Two glasses of orange",
				"juice, please.",
				"duolingo",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			textExtractor := new(AwsRekognition)

			data, err := os.ReadFile(test.FilePath)
			assert.NoError(t, err)

			extractedLines, err := textExtractor.ExtractText(data)
			assert.NoError(t, err)

			expectedJoined := strings.Join(test.ExpectedLines, " / ")
			extractedJoined := strings.Join(extractedLines, " / ")
			assert.Equal(t, expectedJoined, extractedJoined)
		})
	}
}
