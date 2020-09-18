package models

import (
	"database/sql"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

type PreferredTheme struct {
	ID   uuid.UUID `json:"id" db:"id"`
	Name string    `json:"name" db:"name"`

	Theme Theme `json:"theme" db:"-"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (p *PreferredTheme) LoadTheme() {
	p.Theme = Themes[p.Name]
}

func (p *PreferredTheme) Load(tx *pop.Connection) error {
	if err := tx.First(p); err != nil && errors.Cause(err) != sql.ErrNoRows {
		return err
	}

	if p.ID == uuid.Nil {
		p.Name = DayTheme
	}
	p.LoadTheme()
	return nil
}

func (p *PreferredTheme) Save(tx *pop.Connection) error {
	if err := tx.RawQuery("DELETE FROM preferred_themes").Exec(); err != nil {
		return err
	}
	return tx.Save(p)
}
