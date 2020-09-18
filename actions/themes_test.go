package actions_test

import (
	"calculator/models"
	"net/http"
	"net/url"
)

func (as *ActionSuite) Test_SetPreferredTheme() {
	themeNames := []string{
		models.DayTheme,
	}

	for _, themeName := range themeNames {
		res := as.HTML("/theme/").Post(url.Values{
			"Name": []string{themeName},
		})
		as.Equal(http.StatusSeeOther, res.Code)

		preferredTheme := models.PreferredTheme{}

		count, err := as.DB.Count(&preferredTheme)
		as.NoError(err)
		as.Equal(1, count)

		as.NoError(as.DB.First(&preferredTheme))
		as.Equal(themeName, preferredTheme.Name)
	}
}
