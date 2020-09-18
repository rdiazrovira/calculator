package models

const (
	DayTheme   = "day"
	NightTheme = "night"
)

type Theme map[string]string

var Themes = map[string]Theme{
	DayTheme: {
		"primary-background":   "day-primary-background",
		"secondary-background": "day-secondary-background",
		"button":               "day-btn",
		"input":                "day-input",
		"text":                 "day-text",
	},
	NightTheme: {
		"primary-background":   "night-primary-background",
		"secondary-background": "night-secondary-background",
		"button":               "night-btn",
		"input":                "night-input",
		"text":                 "night-text",
	},
}
