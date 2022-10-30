package language_detection

import "github.com/stretchr/testify/mock"

type LanguageDetector interface {
	DetectLanguage([]byte) (string, error)
}

type LanguageDetectorMock struct {
	mock.Mock
}

func (m *LanguageDetectorMock) DetectLanguage(inputBytes []byte) (string, error) {
	args := m.Called(inputBytes)

	return args.Get(0).(string), args.Error(1)
}
