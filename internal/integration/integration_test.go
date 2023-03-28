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
			expectedOutput: "Quest'automobile è come nuova.\nThis automobile is like new.\n",
		},
		{
			name:           "Happy Path With 4 lines - JSON",
			args:           []string{"--json", "../../test/integration.png"},
			expectedError:  "",
			expectedOutput: "[{\"Language\":\"it\",\"Text\":\"Quest'automobile è come nuova.\"},{\"Language\":\"en\",\"Text\":\"This automobile is like new.\"}]",
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

			cmd.Run()

			assert.Equal(t, test.expectedError, errOut.String())
			assert.Equal(t, test.expectedOutput, out.String())
		})
	}
}
