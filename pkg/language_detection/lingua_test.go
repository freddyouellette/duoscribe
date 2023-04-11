package language_detection

import (
	"testing"

	"github.com/pemistahl/lingua-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type LinguaServiceMock struct {
	mock.Mock
}

func (m *LinguaServiceMock) DetectLanguageOf(text string) (lingua.Language, bool) {
	args := m.MethodCalled("DetectLanguageOf", text)
	return args.Get(0).(lingua.Language), args.Bool(1)
}

func TestLingua(t *testing.T) {
	var (
		okInput    = "Lo può ripetere per favore?"
		okLanguage = lingua.Italian
		okOutput   = "it"
	)
	tests := []struct {
		name               string
		serviceMockFactory func() *LinguaServiceMock
		inputString        string
		expectedOutput     string
		expectedErr        error
	}{
		{
			name: "Happy Path",
			serviceMockFactory: func() *LinguaServiceMock {
				m := new(LinguaServiceMock)
				m.On("DetectLanguageOf", okInput).Once().Return(okLanguage, true)
				return m
			},
			inputString:    okInput,
			expectedOutput: okOutput,
			expectedErr:    nil,
		},
		{
			name: "Lingua Failure",
			serviceMockFactory: func() *LinguaServiceMock {
				m := new(LinguaServiceMock)
				m.On("DetectLanguageOf", okInput).Once().Return(lingua.Unknown, false)
				return m
			},
			inputString:    okInput,
			expectedOutput: "",
			expectedErr:    errLinguaFailure,
		},
	}

	for _, test := range tests {
		languageDetector := NewLinguaService(test.serviceMockFactory())

		t.Run(test.name, func(t *testing.T) {
			detectedLanguage, err := languageDetector.DetectLanguage([]byte(test.inputString))
			assert.ErrorIs(t, err, test.expectedErr)
			assert.Equal(t, test.expectedOutput, detectedLanguage)
		})
	}
}
