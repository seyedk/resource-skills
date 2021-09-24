package model

import (
	"database/sql"
	"encoding/json"
)

type Part struct {
	ID        int            `db:"id" json:"id"`
	Name      string         `db:"name" json:"name"`
	Quantity  int            `db:"quantity" json:"quantity"`
	Unit      string         `db:"unit" json:"unit"`
	CreatedAt string         `db:"created_at" json:"-"`
	UpdatedAt string         `db:"updated_at" json:"-"`
	DeletedAt sql.NullString `db:"deleted_at" json:"-"`
}

type Parts []Part

func (p *Parts) ToJSON() ([]byte, error){
	return json.Marshal(p)

}
