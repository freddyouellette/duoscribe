package language_detection

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AwsComprehendServiceMock struct {
	mock.Mock
}

func (m *AwsComprehendServiceMock) DetectDominantLanguage(
	input *comprehend.DetectDominantLanguageInput,
) (*comprehend.DetectDominantLanguageOutput, error) {
	args := m.MethodCalled("DetectDominantLanguage", input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		return args.Get(0).(*comprehend.DetectDominantLanguageOutput), args.Error(1)
	}
}

type languageDetectorTest struct {
	name                        string
	awsComprehendServiceFactory func() *AwsComprehendServiceMock
	inputString                 string
	expectedOutput              string
	expectedErr                 error
}

func TestLanguageDetector(t *testing.T) {
	var (
		okInput           = "Lo può ripetere per favore?"
		okOutput          = "it"
		okComprehendInput = &comprehend.DetectDominantLanguageInput{
			Text: &okInput,
		}
	)
	tests := []languageDetectorTest{
		{
			name: "Happy Path",
			awsComprehendServiceFactory: func() *AwsComprehendServiceMock {
				m := new(AwsComprehendServiceMock)
				output := &comprehend.DetectDominantLanguageOutput{
					Languages: []*comprehend.DominantLanguage{
						{LanguageCode: &okOutput},
					},
				}
				m.On("DetectDominantLanguage", okComprehendInput).Once().Return(output, nil)
				return m
			},
			inputString:    okInput,
			expectedOutput: okOutput,
			expectedErr:    nil,
		},
		{
			name: "AWS Failure",
			awsComprehendServiceFactory: func() *AwsComprehendServiceMock {
				m := new(AwsComprehendServiceMock)
				m.On("DetectDominantLanguage", okComprehendInput).Once().Return(nil, errors.New("aws failure"))
				return m
			},
			inputString:    okInput,
			expectedOutput: "",
			expectedErr:    errAwsComprehendFailure,
		},
	}

	for _, test := range tests {
		languageDetector := NewAwsComprehend(test.awsComprehendServiceFactory())

		t.Run(test.name, func(t *testing.T) {
			detectedLanguage, err := languageDetector.DetectLanguage([]byte(test.inputString))
			assert.ErrorIs(t, err, test.expectedErr)
			assert.Equal(t, test.expectedOutput, detectedLanguage)
		})
	}
}
