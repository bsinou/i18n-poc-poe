package utils

import (
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/nicksnyder/go-i18n/i18n"
	"github.com/nicksnyder/go-i18n/i18n/bundle"
)

type I18nBundle struct {
	bundle.Bundle
}

func NewI18nBundle(box packr.Box) *I18nBundle {
	b := bundle.New()
	B := &I18nBundle{}
	B.Bundle = *b
	B.LoadPackrTranslationFiles(box)
	return B
}

// LoadPackrTranslationFiles loads goi18n translation files from packr boxes
func (b *I18nBundle) LoadPackrTranslationFiles(box packr.Box) {

	files := box.List()
	for _, f := range files {
		if strings.HasSuffix(f, ".json") {
			data := box.Bytes(f)
			b.ParseTranslationFileBytes(f, data)
		}
	}

}

// GetTranslationFunc provides the correct translation func for language or the IdentityFunc
// if language is not supported. Languages can be a list of weighted languages as provided
// in the http header Accept-Language
func (b *I18nBundle) GetTranslationFunc(languages ...string) i18n.TranslateFunc {

	var t bundle.TranslateFunc
	var e error
	if len(languages) > 0 {
		t, e = b.Tfunc(languages[0], "en-US")
	} else {
		t, e = b.Tfunc("en-US")
	}
	if e != nil {
		return i18n.IdentityTfunc()
	}
	return i18n.TranslateFunc(t)

}

// GetDefaultLanguage always returns "en" for this PoC.
func GetDefaultLanguage() string {
	return "en"
}
