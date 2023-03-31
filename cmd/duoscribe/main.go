package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/freddyouellette/duolingo-text-extractor/internal/actions/extract"
	"github.com/freddyouellette/duolingo-text-extractor/internal/output_formatting"
	"github.com/freddyouellette/duolingo-text-extractor/internal/text_cleaning"
	"github.com/freddyouellette/duolingo-text-extractor/internal/text_condensing"
	"github.com/freddyouellette/duolingo-text-extractor/pkg/language_detection"
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

	awsSession, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf(
			"AWS session failed to start. Please check your settings, an AWS environment is required to use this tool: %s",
			err.Error(),
		))
		os.Exit(1)
	}

	textExtractor := text_extraction.NewAwsRekognition(awsSession)
	languageDetector := language_detection.NewAwsComprehend(awsSession)
	textCleaner := text_cleaning.NewTextCleaner()
	textCondenser := text_condensing.NewTextCondenser()

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
