package language_detection

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
)

type AwsComprehend struct{}

func (s *AwsComprehend) DetectLanguage(inputBytes []byte) (string, error) {

	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	service := comprehend.New(session)

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
