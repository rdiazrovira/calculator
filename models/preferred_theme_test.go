package models_test

import (
	"calculator/models"
	"fmt"
)

func (ms *ModelSuite) Test_LoadTheme() {
	tcases := []struct {
		PreferredThemeName string
		Theme              models.Theme
	}{
		{
			PreferredThemeName: "Any theme",
			Theme: models.Theme{
				"primary-background":   "",
				"secondary-background": "",
				"button":               "",
				"input":                "",
				"text":                 "",
			},
		},
		{
			PreferredThemeName: models.DayTheme,
			Theme: models.Theme{
				"primary-background":   "day-primary-background",
				"secondary-background": "day-secondary-background",
				"button":               "day-btn",
				"input":                "day-input",
				"text":                 "day-text",
			},
		},
		{
			PreferredThemeName: models.NightTheme,
			Theme: models.Theme{
				"primary-background":   "night-primary-background",
				"secondary-background": "night-secondary-background",
				"button":               "night-btn",
				"input":                "night-input",
				"text":                 "night-text",
			},
		},
	}

	for index, tcase := range tcases {
		theme := models.PreferredTheme{
			Name: tcase.PreferredThemeName,
		}
		theme.LoadTheme()
		ms.Equal(tcase.Theme["primary-background"], theme.Theme["primary-background"], fmt.Sprintf("index: %v", index))
		ms.Equal(tcase.Theme["secondary-background"], theme.Theme["secondary-background"], fmt.Sprintf("index: %v", index))
		ms.Equal(tcase.Theme["button"], theme.Theme["button"], fmt.Sprintf("index: %v", index))
		ms.Equal(tcase.Theme["input"], theme.Theme["input"], fmt.Sprintf("index: %v", index))
		ms.Equal(tcase.Theme["text"], theme.Theme["text"], fmt.Sprintf("index: %v", index))
	}
}

func (ms *ModelSuite) Test_PreferredTheme_Load() {
	tcases := []struct {
		PreferredThemeName string
		Theme              models.Theme
	}{
		{
			PreferredThemeName: "Any theme",
			Theme: models.Theme{
				"primary-background":   "",
				"secondary-background": "",
				"button":               "",
				"input":                "",
				"text":                 "",
			},
		},
		{
			PreferredThemeName: models.DayTheme,
			Theme: models.Theme{
				"primary-background":   "day-primary-background",
				"secondary-background": "day-secondary-background",
				"button":               "day-btn",
				"input":                "day-input",
				"text":                 "day-text",
			},
		},
		{
			PreferredThemeName: models.NightTheme,
			Theme: models.Theme{
				"primary-background":   "night-primary-background",
				"secondary-background": "night-secondary-background",
				"button":               "night-btn",
				"input":                "night-input",
				"text":                 "night-text",
			},
		},
	}

	for index, tcase := range tcases {
		theme := models.PreferredTheme{
			Name: tcase.PreferredThemeName,
		}
		ms.NoError(ms.DB.Save(&theme), fmt.Sprintf("index: %v", index))
		ms.Equal(tcase.PreferredThemeName, theme.Name, fmt.Sprintf("index: %v", index))

		savedTheme := models.PreferredTheme{}
		ms.NoError(savedTheme.Load(ms.DB), fmt.Sprintf("index: %v", index))
		ms.Equal(tcase.PreferredThemeName, savedTheme.Name, fmt.Sprintf("index: %v", index))
		ms.Equal(tcase.Theme["primary-background"], savedTheme.Theme["primary-background"], fmt.Sprintf("index: %v", index))
		ms.Equal(tcase.Theme["secondary-background"], savedTheme.Theme["secondary-background"], fmt.Sprintf("index: %v", index))
		ms.Equal(tcase.Theme["button"], savedTheme.Theme["button"], fmt.Sprintf("index: %v", index))
		ms.Equal(tcase.Theme["input"], savedTheme.Theme["input"], fmt.Sprintf("index: %v", index))
		ms.Equal(tcase.Theme["text"], savedTheme.Theme["text"], fmt.Sprintf("index: %v", index))

		ms.NoError(ms.DB.RawQuery("DELETE FROM preferred_themes").Exec(), fmt.Sprintf("index: %v", index))
	}
}

func (ms *ModelSuite) Test_PreferredTheme_LoadThemeByDefault() {
	theme := models.PreferredTheme{}
	ms.NoError(theme.Load(ms.DB))

	ms.Equal(models.DayTheme, theme.Name)
	ms.Equal("day-primary-background", theme.Theme["primary-background"])
	ms.Equal("day-btn", theme.Theme["button"])
	ms.Equal("day-input", theme.Theme["input"])
	ms.Equal("day-text", theme.Theme["text"])
}

func (ms *ModelSuite) Test_PreferredTheme_Save() {
	names := []string{
		models.DayTheme,
		models.NightTheme,
	}

	for _, name := range names {
		theme := models.PreferredTheme{
			Name: name,
		}
		ms.NoError(theme.Save(ms.DB))

		count, err := ms.DB.Count(models.PreferredTheme{})
		ms.NoError(err)
		ms.Equal(1, count)

		savedTheme := models.PreferredTheme{}
		ms.NoError(ms.DB.First(&savedTheme))
		ms.Equal(theme.Name, savedTheme.Name)
	}
}
