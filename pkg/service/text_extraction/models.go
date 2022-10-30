package text_extraction

import "github.com/stretchr/testify/mock"

type TextExtractor interface {
	ExtractText([]byte) ([]string, error)
}

type TextExtractorMock struct {
	mock.Mock
}

func (m *TextExtractorMock) ExtractText(inputBytes []byte) ([]string, error) {
	args := m.Called(inputBytes)

	return args.Get(0).([]string), args.Error(1)
}
