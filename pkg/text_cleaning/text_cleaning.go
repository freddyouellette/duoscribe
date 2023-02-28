package text_cleaning

import "strings"

type TextCleaner struct{}

func (s *TextCleaner) CleanText(text string) (string, error) {
	text = strings.ReplaceAll("duolingo", text, "")
	return text, nil
}
