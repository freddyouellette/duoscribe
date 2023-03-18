package text_cleaning

import "strings"

type TextCleaner struct{}

func (s *TextCleaner) CleanText(text string) (string, error) {
	text = strings.ReplaceAll(text, "duolingo", "")
	text = strings.ReplaceAll(text, "Duolingo", "")
	text = strings.Trim(text, " ")
	return text, nil
}
