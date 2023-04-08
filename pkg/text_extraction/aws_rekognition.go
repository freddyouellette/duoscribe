package text_extraction

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
)

// AwsRekognitionService will call an AWS service to extract text from and image
type AwsRekognitionService interface {
	DetectText(ctx context.Context, params *rekognition.DetectTextInput, optFns ...func(*rekognition.Options)) (*rekognition.DetectTextOutput, error)
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
		Image: &types.Image{
			Bytes: inputBytes,
		},
	}

	output, err := a.awsRekognitionService.DetectText(context.Background(), input)
	if err != nil {
		return []string{}, fmt.Errorf("%w: %s", errAwsRekognitionFailure, err)
	}

	var textLines []string
	for _, textDetection := range output.TextDetections {
		if textDetection.Type == types.TextTypesLine {
			textLines = append(textLines, *textDetection.DetectedText)
		}
	}

	return textLines, nil
}
