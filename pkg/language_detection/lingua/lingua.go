package lingua

import (
	"errors"
	"strings"

	"github.com/pemistahl/lingua-go"
)

// ExternalService will call the Lingua package to detect the language of a string
type ExternalService interface {
	DetectLanguageOf(text string) (lingua.Language, bool)
}

// Service can detect the language of a string using the Lingua Service
type Service struct {
	linguaExternalService ExternalService
}

func NewLinguaService(linguaExternalService ExternalService) *Service {
	return &Service{
		linguaExternalService: linguaExternalService,
	}
}

var errLinguaFailure = errors.New("lingua failure")

// DetectLanguage will determine the language of a string and return it.
// The language will be in short form, e.g. "en", "it", "es"
func (s *Service) DetectLanguage(inputBytes []byte) (string, error) {
	inputString := string(inputBytes)
	language, exists := s.linguaExternalService.DetectLanguageOf(inputString)
	if !exists {
		return "", errLinguaFailure
	}

	languageCode := strings.ToLower(language.IsoCode639_1().String())

	return languageCode, nil
}
