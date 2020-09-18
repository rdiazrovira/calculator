package helpers

import (
	"calculator/models"
	"errors"

	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/pop/v5"
)

func ThemeName(ctx hctx.HelperContext) (string, error) {
	tx, ok := ctx.Value("tx").(*pop.Connection)
	if !ok {
		return "", errors.New("no transaction found")
	}

	theme := models.PreferredTheme{}
	err := theme.Load(tx)
	if err != nil {
		return "", err
	}

	return theme.Name, nil
}

func ThemeProperty(property string, ctx hctx.HelperContext) (string, error) {
	tx, ok := ctx.Value("tx").(*pop.Connection)
	if !ok {
		return "", errors.New("no transaction found")
	}

	theme := models.PreferredTheme{}
	err := theme.Load(tx)
	if err != nil {
		return "", err
	}

	return theme.Theme[property], nil
}
