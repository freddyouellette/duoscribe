package text_extraction

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

type AwsRekognition struct {
	awsSession *session.Session
}

func NewAwsRekognition(awsSession *session.Session) *AwsRekognition {
	return &AwsRekognition{
		awsSession: awsSession,
	}
}

func (a *AwsRekognition) ExtractText(inputBytes []byte) ([]string, error) {
	service := rekognition.New(a.awsSession)

	input := &rekognition.DetectTextInput{
		Image: &rekognition.Image{
			Bytes: inputBytes,
		},
	}

	output, err := service.DetectText(input)
	if err != nil {
		return []string{}, err
	}

	var textLines []string
	for _, textDetection := range output.TextDetections {
		if *textDetection.Type == rekognition.TextTypesLine {
			textLines = append(textLines, *textDetection.DetectedText)
		}
	}

	return textLines, nil
}
