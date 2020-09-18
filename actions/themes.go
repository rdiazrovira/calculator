package actions

import (
	"calculator/models"
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/pkg/errors"
)

func SetPreferredTheme(c buffalo.Context) error {
	// Get the pop connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	theme := models.PreferredTheme{}
	if err := c.Bind(&theme); err != nil {
		return errors.Wrap(err, "error by trying to bind the preferred theme")
	}

	if err := theme.Save(tx); err != nil {
		return errors.Wrap(err, "error by trying to save the preferred theme")
	}
	return c.Redirect(http.StatusSeeOther, c.Request().Referer())
}
