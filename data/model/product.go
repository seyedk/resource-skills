package model

import (
	"database/sql"
	"encoding/json"
	"io"
)

type Product struct {
	ID          int            `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	Teaser      string         `db:"teaser" json:"teaser"`
	Description string         `db:"description" json:"description"`
	Price       float64        `db:"price" json:"price"`
	Image       string         `db:"image" json:"image"`
	CreatedAt   string         `db:"created_at" json:"-"`
	UpdatedAt   string         `db:"updated_at" json:"-"`
	DeletedAt   sql.NullString `db:"deleted_at" json:"-"`
	Parts       []ProductPart  `json:"parts"`
}

type ProductPart struct {
	ID        int            `db:"id" json:"-"`
	ProductID int            `db:"product_id" json:"-"`
	PartID    int            `db:"part_id" json:"ingredient_id"`
	CreatedAt string         `db:"created_at" json:"-"`
	UpdatedAt string         `db:"updated_at" json:"-"`
	DeletedAt sql.NullString `db:"deleted_at" json:"-"`
}

type Products []Product

func (p *Products) FromJSON(data io.Reader) error {
	de := json.NewDecoder(data)
	return de.Decode(p)

}

func (p *Products) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Product) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}
func (p *Product) FromJSON(data io.Reader) error {
	de := json.NewDecoder(data)
	return de.Decode(p)
}

func (p *ProductPart) ToJSON() ([]byte, error) {

	return json.Marshal(p)
}
