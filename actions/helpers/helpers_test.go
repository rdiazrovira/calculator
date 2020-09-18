package helpers_test

import (
	"calculator/actions"
	"calculator/actions/helpers"
	"calculator/models"
	"testing"

	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush/v4"
	"github.com/gobuffalo/suite/v3"
)

type HelperSuite struct {
	*suite.Action
}

func Test_HelperSuite(t *testing.T) {
	action, err := suite.NewActionWithFixtures(actions.App(), packr.New("Test_HelperSuite", "../../fixtures"))
	if err != nil {
		t.Fatal(err)
	}

	hs := &HelperSuite{
		Action: action,
	}
	suite.Run(t, hs)
}

func (hs *HelperSuite) Test_ThemeName() {
	ctx := plush.HelperContext{Context: plush.NewContext()}
	ctx.Set("tx", hs.DB)

	themeName, err := helpers.ThemeName(ctx)
	hs.NoError(err)
	hs.Equal(models.DayTheme, themeName)

	theme := models.PreferredTheme{
		Name: models.NightTheme,
	}
	err = theme.Save(hs.DB)
	hs.NoError(err)

	themeName, err = helpers.ThemeName(ctx)
	hs.NoError(err)
	hs.Equal(models.NightTheme, themeName)
}

func (hs *HelperSuite) Test_ThemeProperty() {
	ctx := plush.HelperContext{Context: plush.NewContext()}
	ctx.Set("tx", hs.DB)

	for name, properties := range models.Themes {
		theme := models.PreferredTheme{
			Name: name,
		}
		err := theme.Save(hs.DB)
		hs.NoError(err)

		for propName, propValue := range properties {
			themeProperty, err := helpers.ThemeProperty(propName, ctx)
			hs.NoError(err)
			hs.Equal(propValue, themeProperty)
		}
	}
}
