package text_extractor

import "github.com/stretchr/testify/mock"

type TextExtractorService interface {
	ExtractText([]byte) ([]string, error)
}

type TextExtractorServiceMock struct {
	mock.Mock
}

func (m *TextExtractorServiceMock) ExtractText(inputBytes []byte) ([]string, error) {
	args := m.Called(inputBytes)

	return args.Get(0).([]string), args.Error(1)
}
