package text_cleaning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type textCleaningTest struct {
	Name           string
	Input          string
	ExpectedOutput string
	ExpectedError  error
}

func TestTextCleaning(t *testing.T) {
	tests := []textCleaningTest{
		{
			Name:           "No invalid text",
			Input:          "no invalid text",
			ExpectedOutput: "no invalid text",
			ExpectedError:  nil,
		},
		{
			Name:           "With lowercase duolingo",
			Input:          "With lowercase duolingo",
			ExpectedOutput: "With lowercase",
			ExpectedError:  nil,
		},
		{
			Name:           "With Uppercase Duolingo",
			Input:          "With Uppercase Duolingo",
			ExpectedOutput: "With Uppercase",
			ExpectedError:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			textCleaner := new(TextCleaner)
			actualOutput, err := textCleaner.CleanText(test.Input)

			if test.ExpectedError != nil {
				assert.ErrorAs(t, err, test.ExpectedError)
			} else {
				assert.Equal(t, test.ExpectedOutput, actualOutput)
			}
		})
	}
}
