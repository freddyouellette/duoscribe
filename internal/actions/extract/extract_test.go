package extract

import (
	"errors"
	"testing"

	"github.com/freddyouellette/duolingo-text-extractor/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TextExtractorMock struct {
	mock.Mock
}

func (mock *TextExtractorMock) ExtractText(inputBytes []byte) ([]string, error) {
	args := mock.Called(inputBytes)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
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
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Text), args.Error(1)
}

type OutputterMock struct {
	mock.Mock
}

func (mock *OutputterMock) Render(output []models.Text) (string, error) {
	args := mock.Called(output)
	return args.String(0), args.Error(1)
}

type extractTest struct {
	name             string
	inputImageBytes  []byte
	textExtractor    func() *TextExtractorMock
	textCleaner      func() *TextCleanerMock
	languageDetector func() *LanguageDetectorMock
	textCondenser    func() *TextCondenserMock
	outputter        func() *OutputterMock
	expectedOutput   string
	expectedError    error
}

func TestLanguageDetector(t *testing.T) {
	testImageBytes := []byte("testImageBytes")
	testExtractedTextLines := []string{"Test Extracted Text"}
	testDetectedLanguage := "en"
	testCleanedText := "Test Cleaned Text"
	testPreCondensedTexts := []models.Text{{Language: "en", Text: "Test Cleaned Text"}}
	testCondensedTexts := []models.Text{{Language: "en", Text: "Test Condensed Text"}}
	testOutputtedText := "Test Outputted Text"

	okTextExtractorFactory := func() *TextExtractorMock {
		mock := new(TextExtractorMock)
		mock.On("ExtractText", testImageBytes).Once().Return(testExtractedTextLines, nil)
		return mock
	}
	okTextCleanerFactory := func() *TextCleanerMock {
		mock := new(TextCleanerMock)
		mock.On("CleanText", testExtractedTextLines[0]).Once().Return(testCleanedText, nil)
		return mock
	}
	okLanguageDetectorFactory := func() *LanguageDetectorMock {
		mock := new(LanguageDetectorMock)
		mock.On("DetectLanguage", []byte(testCleanedText)).Once().Return(testDetectedLanguage, nil)
		return mock
	}
	okTextCondenserFactory := func() *TextCondenserMock {
		mock := new(TextCondenserMock)
		mock.On("Condense", testPreCondensedTexts).Once().Return(testCondensedTexts, nil)
		return mock
	}
	okOutputterFactory := func() *OutputterMock {
		mock := new(OutputterMock)
		mock.On("Render", testCondensedTexts).Once().Return(testOutputtedText, nil)
		return mock
	}
	noopTextCleanerFactory := func() *TextCleanerMock { return new(TextCleanerMock) }
	noopLanguageDetectorFactory := func() *LanguageDetectorMock { return new(LanguageDetectorMock) }
	noopTextCondenserFactory := func() *TextCondenserMock { return new(TextCondenserMock) }
	noopOutputterFactory := func() *OutputterMock { return new(OutputterMock) }

	tests := []extractTest{
		{
			name:             "Successful",
			inputImageBytes:  testImageBytes,
			textExtractor:    okTextExtractorFactory,
			textCleaner:      okTextCleanerFactory,
			languageDetector: okLanguageDetectorFactory,
			textCondenser:    okTextCondenserFactory,
			outputter:        okOutputterFactory,
			expectedOutput:   testOutputtedText,
			expectedError:    nil,
		},
		{
			name:            "Text Extraction Failed",
			inputImageBytes: testImageBytes,
			textExtractor: func() *TextExtractorMock {
				mock := new(TextExtractorMock)
				mock.On("ExtractText", testImageBytes).Once().Return(nil, errors.New("text extraction failed"))
				return mock
			},
			textCleaner:      noopTextCleanerFactory,
			languageDetector: noopLanguageDetectorFactory,
			textCondenser:    noopTextCondenserFactory,
			outputter:        noopOutputterFactory,
			expectedOutput:   "",
			expectedError:    errTextExtractor,
		},
		{
			name:            "Text Cleaning Failed",
			inputImageBytes: testImageBytes,
			textExtractor:   okTextExtractorFactory,
			textCleaner: func() *TextCleanerMock {
				mock := new(TextCleanerMock)
				mock.On("CleanText", testExtractedTextLines[0]).Once().Return("", errors.New("text cleaning failure"))
				return mock
			},
			languageDetector: noopLanguageDetectorFactory,
			textCondenser:    noopTextCondenserFactory,
			outputter:        noopOutputterFactory,
			expectedOutput:   "",
			expectedError:    errTextCleaner,
		},
		{
			name:            "Language Detection Failed",
			inputImageBytes: testImageBytes,
			textExtractor:   okTextExtractorFactory,
			textCleaner:     okTextCleanerFactory,
			languageDetector: func() *LanguageDetectorMock {
				mock := new(LanguageDetectorMock)
				mock.On("DetectLanguage", []byte(testCleanedText)).Once().Return("", errors.New("language detection failed"))
				return mock
			},
			textCondenser:  noopTextCondenserFactory,
			outputter:      noopOutputterFactory,
			expectedOutput: "",
			expectedError:  errLanguageDetector,
		},
		{
			name:             "Text Condensation Failed",
			inputImageBytes:  testImageBytes,
			textExtractor:    okTextExtractorFactory,
			textCleaner:      okTextCleanerFactory,
			languageDetector: okLanguageDetectorFactory,
			textCondenser: func() *TextCondenserMock {
				mock := new(TextCondenserMock)
				mock.On("Condense", testPreCondensedTexts).Once().Return(nil, errors.New("text condensation failed"))
				return mock
			},
			outputter:      noopOutputterFactory,
			expectedOutput: "",
			expectedError:  errTextCondenser,
		},
		{
			name:             "Outputter Failed",
			inputImageBytes:  testImageBytes,
			textExtractor:    okTextExtractorFactory,
			textCleaner:      okTextCleanerFactory,
			languageDetector: okLanguageDetectorFactory,
			textCondenser:    okTextCondenserFactory,
			outputter: func() *OutputterMock {
				mock := new(OutputterMock)
				mock.On("Render", testCondensedTexts).Once().Return("", errors.New("outputter failed"))
				return mock
			},
			expectedOutput: "",
			expectedError:  errOutputter,
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

			actualOutput, actualErr := extractorAction.Extract(test.inputImageBytes)
			assert.ErrorIs(t, actualErr, test.expectedError)
			assert.Equal(t, test.expectedOutput, actualOutput)

			textExtractorMock.AssertExpectations(t)
			languageDetectorMock.AssertExpectations(t)
			textCleanerMock.AssertExpectations(t)
			textCondenserMock.AssertExpectations(t)
			outputterMock.AssertExpectations(t)
		})
	}
}
