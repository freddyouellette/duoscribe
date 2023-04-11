package aws_comprehend

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/comprehend"
)

// ExternalService will call the AWS service to detect the language of a string.
type ExternalService interface {
	DetectDominantLanguage(ctx context.Context, params *comprehend.DetectDominantLanguageInput, optFns ...func(*comprehend.Options)) (*comprehend.DetectDominantLanguageOutput, error)
}

// Service can detect the language of a string using an AWS service.
type Service struct {
	awsComprehendService ExternalService
}

func NewAwsComprehend(awsComprehendService ExternalService) *Service {
	return &Service{
		awsComprehendService: awsComprehendService,
	}
}

var errAwsComprehendFailure = errors.New("aws comprehend failure")

// DetectLanguage will determine the language of a string and return it.
// The language will be in short form, e.g. "en", "it", "es"
func (s *Service) DetectLanguage(inputBytes []byte) (string, error) {
	inputString := string(inputBytes)
	input := &comprehend.DetectDominantLanguageInput{
		Text: &inputString,
	}

	output, err := s.awsComprehendService.DetectDominantLanguage(context.Background(), input)
	if err != nil {
		return "", fmt.Errorf("%w: %s", errAwsComprehendFailure, err)
	}

	return *output.Languages[0].LanguageCode, nil
}
