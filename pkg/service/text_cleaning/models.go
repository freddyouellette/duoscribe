package text_cleaning

type TextCleaner interface {
	ExtractTranslations(inputBytes []byte) (originalString string, translatedString string, err error)
}

type TextOrganizerOutput struct {
	Texts []Text
}

type Text struct {
	Language string
	Text     string
}
