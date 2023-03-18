package extract

import (
	"testing"

	"github.com/freddyouellette/duolingo-text-extractor/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ExtractTest struct {
	name             string
	inputImageBytes  []byte
	textExtractor    func() *TextExtractorMock
	textCleaner      func() *TextCleanerMock
	languageDetector func() *LanguageDetectorMock
	textCondenser    func() *TextCondenserMock
	outputter        func() *OutputterMock
	expectedError    error
}

type TextExtractorMock struct {
	mock.Mock
}

func (mock *TextExtractorMock) ExtractText(inputBytes []byte) ([]string, error) {
	args := mock.Called(inputBytes)
	return args.Get(0).([]string), args.Error(1)
}

type TextCleanerMock struct {
	mock.Mock
}

func (mock *TextCleanerMock) CleanText(text string) (string, error) {
	args := mock.Called(text)
	return args.String(0), args.Error(1)
}

type LanguageDetectorMock struct {
	mock.Mock
}

func (mock *LanguageDetectorMock) DetectLanguage(inputBytes []byte) (string, error) {
	args := mock.Called(inputBytes)
	return args.String(0), args.Error(1)
}

type TextCondenserMock struct {
	mock.Mock
}

func (mock *TextCondenserMock) Condense(texts []models.Text) ([]models.Text, error) {
	args := mock.Called(texts)
	return args.Get(0).([]models.Text), args.Error(1)
}

type OutputterMock struct {
	mock.Mock
}

func (mock *OutputterMock) Render(output []models.Text) (string, error) {
	args := mock.Called(output)
	return args.String(0), args.Error(1)
}

func TestLanguageDetector(t *testing.T) {
	testImageBytes := []byte("testImageBytes")
	testExtractedTextLines := []string{"Test Extracted Text"}
	testDetectedLanguage := "en"
	testCleanedText := "Test Cleaned Text"
	testPreCondensedTexts := []models.Text{{Language: "en", Text: "Test Cleaned Text"}}
	testCondensedTexts := []models.Text{{Language: "en", Text: "Test Condensed Text"}}
	testOutputtedText := "Test Outputted Text"
	tests := []ExtractTest{
		{
			name:            "Successful",
			inputImageBytes: testImageBytes,
			textExtractor: func() *TextExtractorMock {
				mock := new(TextExtractorMock)
				mock.On("ExtractText", testImageBytes).Return(testExtractedTextLines, nil)
				return mock
			},
			textCleaner: func() *TextCleanerMock {
				mock := new(TextCleanerMock)
				mock.On("CleanText", testExtractedTextLines[0]).Return(testCleanedText, nil)
				return mock
			},
			languageDetector: func() *LanguageDetectorMock {
				mock := new(LanguageDetectorMock)
				mock.On("DetectLanguage", []byte(testCleanedText)).Return(testDetectedLanguage, nil)
				return mock
			},
			textCondenser: func() *TextCondenserMock {
				mock := new(TextCondenserMock)
				mock.On("Condense", testPreCondensedTexts).Return(testCondensedTexts, nil)
				return mock
			},
			outputter: func() *OutputterMock {
				mock := new(OutputterMock)
				mock.On("Render", testCondensedTexts).Return(testOutputtedText, nil)
				return mock
			},
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			textExtractorMock := test.textExtractor()
			languageDetectorMock := test.languageDetector()
			textCleanerMock := test.textCleaner()
			textCondenserMock := test.textCondenser()
			outputterMock := test.outputter()

			extractorAction := &Action{
				TextExtractor:    textExtractorMock,
				LanguageDetector: languageDetectorMock,
				TextCleaner:      textCleanerMock,
				TextCondenser:    textCondenserMock,
				Outputter:        outputterMock,
			}

			err := extractorAction.Extract(test.inputImageBytes)

			textExtractorMock.AssertExpectations(t)
			languageDetectorMock.AssertExpectations(t)
			textCleanerMock.AssertExpectations(t)
			textCondenserMock.AssertExpectations(t)
			outputterMock.AssertExpectations(t)

			assert.ErrorIs(t, err, test.expectedError)
		})
	}
}
