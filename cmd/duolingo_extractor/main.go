package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/freddyouellette/duolingo-text-extractor/pkg/actions/extract"
	"github.com/freddyouellette/duolingo-text-extractor/pkg/language_detection"
	"github.com/freddyouellette/duolingo-text-extractor/pkg/output_formatting"
	"github.com/freddyouellette/duolingo-text-extractor/pkg/text_cleaning"
	"github.com/freddyouellette/duolingo-text-extractor/pkg/text_condensing"
	"github.com/freddyouellette/duolingo-text-extractor/pkg/text_extraction"
)

func main() {
	var wantJson bool
	for _, arg := range os.Args {
		if arg == "--json" {
			wantJson = true
		}
	}

	filePath := os.Args[len(os.Args)-1]
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("File %s cannot be read.", filePath))
		os.Exit(1)
	}

	textExtractor := &text_extraction.AwsRekognition{}
	languageDetector := &language_detection.AwsComprehend{}
	textCleaner := &text_cleaning.TextCleaner{}
	textCondenser := &text_condensing.TextCondenser{}

	var outputter extract.Outputter
	if wantJson {
		outputter = &output_formatting.JsonOutputter{}
	} else {
		outputter = &output_formatting.TextOutputter{}

	}

	action := extract.Action{
		TextExtractor:    textExtractor,
		LanguageDetector: languageDetector,
		TextCleaner:      textCleaner,
		TextCondenser:    textCondenser,
		Outputter:        outputter,
	}

	err = action.Extract(fileBytes)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
}
