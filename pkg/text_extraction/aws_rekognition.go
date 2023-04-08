package text_extraction

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/service/rekognition"
)

type AwsRekognitionService interface {
	DetectText(input *rekognition.DetectTextInput) (*rekognition.DetectTextOutput, error)
}

// AwsRekognition can extract text from an image using an AWS service.
type AwsRekognition struct {
	awsRekognitionService AwsRekognitionService
}

func NewAwsRekognition(awsRekognitionService AwsRekognitionService) *AwsRekognition {
	return &AwsRekognition{
		awsRekognitionService: awsRekognitionService,
	}
}

var errAwsRekognitionFailure = errors.New("aws rekognition failure")

// ExtractText will return an array of strings that was extracted from the image given.
func (a *AwsRekognition) ExtractText(inputBytes []byte) ([]string, error) {
	input := &rekognition.DetectTextInput{
		Image: &rekognition.Image{
			Bytes: inputBytes,
		},
	}

	output, err := a.awsRekognitionService.DetectText(input)
	if err != nil {
		return []string{}, fmt.Errorf("%w: %s", errAwsRekognitionFailure, err)
	}

	var textLines []string
	for _, textDetection := range output.TextDetections {
		if *textDetection.Type == rekognition.TextTypesLine {
			textLines = append(textLines, *textDetection.DetectedText)
		}
	}

	return textLines, nil
}
