package actions

import (
	"calculator/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

func CalculateShow(c buffalo.Context) error {
	c.Set("calculator", models.Calculator{})
	return c.Render(http.StatusOK, r.HTML("/calculator/index.plush.html"))
}

func Calculate(c buffalo.Context) error {
	calculator := models.Calculator{}
	if err := c.Bind(&calculator); err != nil {
		return errors.Wrap(err, "error by trying to bind a calculator")
	}

	if err := calculator.Calculate(); err != nil {
		c.Logger().Debug("error: ", err)
	}
	c.Set("calculator", calculator)
	return c.Render(http.StatusOK, r.HTML("/calculator/index.plush.html"))
}
