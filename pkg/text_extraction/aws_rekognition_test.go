package text_extraction

import (
	"errors"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AwsRekognitionServiceMock struct {
	mock.Mock
}

func (m *AwsRekognitionServiceMock) DetectText(input *rekognition.DetectTextInput) (*rekognition.DetectTextOutput, error) {
	args := m.MethodCalled("DetectText", input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		return args.Get(0).(*rekognition.DetectTextOutput), args.Error(1)
	}
}

type rekognitionExtractFromFileTest struct {
	Name                         string
	AwsRekognitionServiceFactory func() *AwsRekognitionServiceMock
	FilePath                     string
	ExpectedOutput               []string
	ExpectedError                error
}

func TestRekognitionExtractFromFile(t *testing.T) {
	var (
		okImageBytes       = []byte("ok-image-bytes")
		okRekognitionInput = &rekognition.DetectTextInput{
			Image: &rekognition.Image{
				Bytes: okImageBytes,
			},
		}
		okText1    = "ok-test-1"
		okText2    = "ok-test-2"
		okText3    = "ok-test-3"
		okTextType = rekognition.TextTypesLine
	)

	tests := []rekognitionExtractFromFileTest{
		{
			Name: "Happy Path",
			AwsRekognitionServiceFactory: func() *AwsRekognitionServiceMock {
				m := new(AwsRekognitionServiceMock)
				output := &rekognition.DetectTextOutput{
					TextDetections: []*rekognition.TextDetection{
						{
							DetectedText: &okText1,
							Type:         &okTextType,
						},
						{
							DetectedText: &okText2,
							Type:         &okTextType,
						},
						{
							DetectedText: &okText3,
							Type:         &okTextType,
						},
					},
				}
				m.On("DetectText", okRekognitionInput).Once().Return(output, nil)
				return m
			},
			FilePath:       "../../test/2_lines.jpg",
			ExpectedOutput: []string{okText1, okText2, okText3},
			ExpectedError:  nil,
		},
		{
			Name: "AWS Failure",
			AwsRekognitionServiceFactory: func() *AwsRekognitionServiceMock {
				m := new(AwsRekognitionServiceMock)
				m.On("DetectText", okRekognitionInput).Once().Return(nil, errors.New("aws failure"))
				return m
			},
			FilePath:       "../../test/2_lines.jpg",
			ExpectedOutput: nil,
			ExpectedError:  errAwsRekognitionFailure,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			textExtractor := NewAwsRekognition(test.AwsRekognitionServiceFactory())

			extractedLines, err := textExtractor.ExtractText(okImageBytes)
			assert.ErrorIs(t, err, test.ExpectedError)

			expectedJoined := strings.Join(test.ExpectedOutput, " / ")
			extractedJoined := strings.Join(extractedLines, " / ")
			assert.Equal(t, expectedJoined, extractedJoined)
		})
	}
}
