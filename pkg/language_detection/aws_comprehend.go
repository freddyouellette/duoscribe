package language_detection

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
)

// AwsComprehend can detect the language of a string using an AWS service.
type AwsComprehend struct {
	awsSession *session.Session
}

func NewAwsComprehend(awsSession *session.Session) *AwsComprehend {
	return &AwsComprehend{
		awsSession: awsSession,
	}
}

// DetectLanguage will determine the language of a string and return it.
// The language will be in short form, e.g. "en", "it", "es"
func (s *AwsComprehend) DetectLanguage(inputBytes []byte) (string, error) {
	service := comprehend.New(s.awsSession)

	inputString := string(inputBytes)
	input := &comprehend.DetectDominantLanguageInput{
		Text: &inputString,
	}

	output, err := service.DetectDominantLanguage(input)
	if err != nil {
		return "", err
	}

	return *output.Languages[0].LanguageCode, nil
}
