package output_formatting

import "github.com/freddyouellette/duolingo-text-extractor/pkg/service/text_cleaning"

type OutputFormatter interface {
	Render(text_cleaning.TextCleanerOutput) (string, error)
}
