package main

import (
	"bytes"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

type textCleaningTest struct {
	name           string
	args           []string
	expectedError  string
	expectedOutput string
}

func TestIntegration(t *testing.T) {
	tests := []textCleaningTest{
		{
			name:           "Happy Path With 4 lines",
			args:           []string{"../../test/integration.png"},
			expectedError:  "",
			expectedOutput: "This automobile is like new.\nQuest'automobile è come nuova.\n",
		},
		{
			name:           "Happy Path With 4 lines - JSON",
			args:           []string{"--json", "../../test/integration.png"},
			expectedError:  "",
			expectedOutput: "[{\"Language\":\"en\",\"Text\":\"This automobile is like new.\"},{\"Language\":\"it\",\"Text\":\"Quest'automobile è come nuova.\"}]",
		},
		{
			name:           "Invalid Path",
			args:           []string{"../../test/invalid_path.png"},
			expectedError:  "File ../../test/invalid_path.png cannot be read.",
			expectedOutput: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := exec.Command("../../bin/duoscribe", test.args...)

			var out, errOut bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &errOut

			err := cmd.Run()
			if err != nil {
				assert.Equal(t, test.expectedError, errOut.String())
			} else {
				assert.Equal(t, test.expectedOutput, out.String())
			}
		})
	}
}
