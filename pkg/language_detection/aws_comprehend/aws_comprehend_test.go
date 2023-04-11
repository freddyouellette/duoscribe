package aws_comprehend

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/comprehend"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ExternalServiceMock struct {
	mock.Mock
}

func (m *ExternalServiceMock) DetectDominantLanguage(ctx context.Context, params *comprehend.DetectDominantLanguageInput, optFns ...func(*comprehend.Options)) (*comprehend.DetectDominantLanguageOutput, error) {
	args := m.MethodCalled("DetectDominantLanguage", ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*comprehend.DetectDominantLanguageOutput), args.Error(1)
}

type languageDetectorTest struct {
	name                        string
	awsComprehendServiceFactory func() *ExternalServiceMock
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
			awsComprehendServiceFactory: func() *ExternalServiceMock {
				m := new(ExternalServiceMock)
				output := &comprehend.DetectDominantLanguageOutput{
					Languages: []types.DominantLanguage{
						{LanguageCode: &okOutput},
					},
				}
				m.On("DetectDominantLanguage", mock.Anything, okComprehendInput, mock.Anything).Once().Return(output, nil)
				return m
			},
			inputString:    okInput,
			expectedOutput: okOutput,
			expectedErr:    nil,
		},
		{
			name: "AWS Failure",
			awsComprehendServiceFactory: func() *ExternalServiceMock {
				m := new(ExternalServiceMock)
				m.On("DetectDominantLanguage", mock.Anything, okComprehendInput, mock.Anything).Once().Return(nil, errors.New("aws failure"))
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
