package language_detection

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/service/comprehend"
)

// AwsComprehendService will call the AWS service to detect the language of a string.
type AwsComprehendService interface {
	DetectDominantLanguage(input *comprehend.DetectDominantLanguageInput) (*comprehend.DetectDominantLanguageOutput, error)
}

// AwsComprehend can detect the language of a string using an AWS service.
type AwsComprehend struct {
	awsComprehendService AwsComprehendService
}

func NewAwsComprehend(awsComprehendService AwsComprehendService) *AwsComprehend {
	return &AwsComprehend{
		awsComprehendService: awsComprehendService,
	}
}

var errAwsComprehendFailure = errors.New("aws comprehend failure")

// DetectLanguage will determine the language of a string and return it.
// The language will be in short form, e.g. "en", "it", "es"
func (s *AwsComprehend) DetectLanguage(inputBytes []byte) (string, error) {
	inputString := string(inputBytes)
	input := &comprehend.DetectDominantLanguageInput{
		Text: &inputString,
	}

	output, err := s.awsComprehendService.DetectDominantLanguage(input)
	if err != nil {
		return "", fmt.Errorf("%w: %s", errAwsComprehendFailure, err)
	}

	return *output.Languages[0].LanguageCode, nil
}
