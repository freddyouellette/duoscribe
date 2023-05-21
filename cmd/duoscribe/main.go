package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/freddyouellette/duolingo-text-extractor/internal/actions/extract"
	"github.com/freddyouellette/duolingo-text-extractor/internal/output_formatting"
	"github.com/freddyouellette/duolingo-text-extractor/internal/text_cleaning"
	"github.com/freddyouellette/duolingo-text-extractor/internal/text_condensing"
	lingua_service "github.com/freddyouellette/duolingo-text-extractor/pkg/language_detection/lingua"
	"github.com/freddyouellette/duolingo-text-extractor/pkg/text_extraction"
	"github.com/pemistahl/lingua-go"
)

var (
	ErrNoArgs           = errors.New("at least one arg is required: the image file to be read")
	ErrFileCannotBeRead = errors.New("file cannot be read")
	ErrAwsConfig        = errors.New("AWS session failed to start. Please check your settings, an AWS environment is required to use this tool")
)

func main() {
	var wantJson bool
	for _, arg := range os.Args {
		if arg == "--json" {
			wantJson = true
		}
	}

	if len(os.Args) == 1 {
		os.Stderr.WriteString(ErrNoArgs.Error())
	}

	filePath := os.Args[len(os.Args)-1]
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		os.Stderr.WriteString(fmt.Errorf("%w: %s", ErrFileCannotBeRead, filePath).Error())
		os.Exit(1)
	}

	awsConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		os.Stderr.WriteString(fmt.Errorf("%w: %s", ErrAwsConfig, err).Error())
		os.Exit(1)
	}

	awsRekognitionService := rekognition.NewFromConfig(awsConfig)
	textExtractor := text_extraction.NewAwsRekognition(awsRekognitionService)
	linguaExternalService := lingua.NewLanguageDetectorBuilder().FromAllLanguages().Build()
	languageDetector := lingua_service.NewLinguaService(linguaExternalService)
	textCleaner := text_cleaning.NewTextCleaner()
	textCondenser := text_condensing.NewTextCondenser()

	var outputter extract.Outputter
	if wantJson {
		outputter = output_formatting.NewJsonOutputter(json.Marshal)
	} else {
		outputter = output_formatting.NewTextOutputter()
	}

	action := extract.Action{
		TextExtractor:    textExtractor,
		LanguageDetector: languageDetector,
		TextCleaner:      textCleaner,
		TextCondenser:    textCondenser,
		Outputter:        outputter,
	}

	output, err := action.Extract(fileBytes)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
	fmt.Print(output)
}
