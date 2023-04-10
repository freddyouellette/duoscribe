package main

import (
	"bytes"
	"os"
	"os/exec"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

type textCleaningTest struct {
	name               string
	args               []string
	env                []string
	expectedErrorRegex string
	expectedOutput     string
}

func TestIntegration(t *testing.T) {
	tests := []textCleaningTest{
		{
			name:               "Happy Path With 4 lines",
			args:               []string{"../../test/integration.png"},
			env:                nil,
			expectedErrorRegex: "",
			expectedOutput:     "This automobile is like new.\nQuest'automobile è come nuova.\n",
		},
		{
			name:               "Happy Path With 4 lines - JSON",
			args:               []string{"--json", "../../test/integration.png"},
			env:                nil,
			expectedErrorRegex: "",
			expectedOutput:     "[{\"Language\":\"en\",\"Text\":\"This automobile is like new.\"},{\"Language\":\"it\",\"Text\":\"Quest'automobile è come nuova.\"}]",
		},
		{
			name:               "Invalid Path",
			args:               []string{"../../test/invalid_path.png"},
			env:                nil,
			expectedErrorRegex: `File .* cannot be read\.`,
			expectedOutput:     "",
		},
		{
			name:               "Invalid AWS",
			args:               []string{"../../test/integration.png"},
			env:                []string{"AWS_ACCESS_KEY_ID=xxxxx"},
			expectedErrorRegex: `.* The security token included in the request is invalid\.`,
			expectedOutput:     "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := exec.Command("../../bin/duoscribe", test.args...)
			cmd.Env = append(os.Environ(), test.env...)

			var out, errOut bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &errOut

			err := cmd.Run()
			if err != nil {
				assert.Regexp(t, regexp.MustCompile(test.expectedErrorRegex), errOut.String())
			} else {
				assert.Equal(t, test.expectedOutput, out.String())
			}
		})
	}
}
