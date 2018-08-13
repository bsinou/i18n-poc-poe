package main

import (
	"context"

	"github.com/bsinou/i18n-poc-poe/goui"
)

func main() {
	goui.LogTranslatedMessage(context.Background())
}
