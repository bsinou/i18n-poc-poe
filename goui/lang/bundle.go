// Package lang provides auth-related i18n strings
package lang

import (
	"github.com/gobuffalo/packr"

	"github.com/bsinou/i18n-poc-poe/utils"
)

var (
	bundle *utils.I18nBundle
)

func Bundle() *utils.I18nBundle {
	if bundle == nil {
		bundle = utils.NewI18nBundle(packr.NewBox("../../goui/lang/box"))
	}
	return bundle
}
