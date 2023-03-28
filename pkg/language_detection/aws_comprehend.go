package language_detection

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
)

type AwsComprehend struct {
	awsSession *session.Session
}

func NewAwsComprehend(awsSession *session.Session) *AwsComprehend {
	return &AwsComprehend{
		awsSession: awsSession,
	}
}

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
