package main

import (
	"log"

	"context"

	"cloud.google.com/go/translate"
	"google.golang.org/api/option"

	"golang.org/x/text/language"
)

type translation struct {
	text   string
	result string
	lang   language.Tag
	target language.Tag
	err    error
}

func newTranslation() *translation {

	t := new(translation)

	t.target = language.English

	return t
}

func (t *translation) clear() {

	t.text = ""
	t.result = ""
	t.lang = language.Und
	t.target = language.English
	t.err = nil

}

func (t *translation) do(text string) error {

	t.clear()

	t.text = text

	t.lang, t.err = detectLanguage(t.text)
	if t.err != nil {
		return t.err
	}

	if t.lang == language.English {
		t.target = language.Japanese
	}

	t.result, t.err = doTranslation(t.text, t.target)
	if t.err != nil {
		return t.err
	}

	return nil
}

func detectLanguage(text string) (language.Tag, error) {

	opt := option.WithCredentialsFile(config.getCredentialsFilePath())
	ctx := context.Background()
	client, err := translate.NewClient(ctx, opt)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return language.Und, err
	}
	defer client.Close()

	lang, err := client.DetectLanguage(ctx, []string{text})
	if err != nil {
		log.Fatalf("Failed to detect language: %v", err)
		return language.Und, err
	}

	return lang[0][0].Language, nil
}

func doTranslation(text string, target language.Tag) (string, error) {

	ctx := context.Background()
	result := ""

	opt := option.WithCredentialsFile(config.getCredentialsFilePath())
	client, err := translate.NewClient(ctx, opt)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return result, err
	}
	defer client.Close()

	translations, err := client.Translate(ctx, []string{text}, target, nil)
	if err != nil {
		log.Fatalf("Failed to translate text: %v", err)
		return result, err
	}
	result = translations[0].Text

	return result, err
}
