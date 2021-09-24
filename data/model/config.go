package model

import (
	"database/sql"
	"encoding/json"
	"io"
)

type Config struct {
	ID        int            `db:"id" json:"id"`
	Name      string         `db:"name" json:"name"`
	Value     string         `db:"value" json:"value"`
	CreatedAt string         `db:"created_at" json:"-"`
	UpdatedAt string         `db:"updated_at" json:"-"`
	DeletedAt sql.NullString `db:"deleted_at" json:"-"`
}

type Configs []Config

func (c *Configs) ToJSON() ([]byte, error) {

	return json.Marshal(c)

}

func (c *Configs) FromJSON(data io.Reader) error {
	de := json.NewDecoder(data)
	return de.Decode(c)
}
