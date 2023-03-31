package text_cleaning

import "strings"

// TextCleaner will clean text of extraneous info, like the logo "duolingo"
type TextCleaner struct{}

func NewTextCleaner() *TextCleaner {
	return &TextCleaner{}
}

// CleanText strips a string of all words that don't belong, particularly the logo text "duolingo"
func (s *TextCleaner) CleanText(text string) (string, error) {
	text = strings.ReplaceAll(text, "duolingo", "")
	text = strings.ReplaceAll(text, "Duolingo", "")
	text = strings.Trim(text, " ")
	return text, nil
}
