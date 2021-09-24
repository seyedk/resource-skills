package model

import (
	"database/sql"
	"encoding/json"
	"io"
)

type Orders []Order

type Order struct {
	ID        int            `db:"id" json:"id,omitempty"`
	UserID    int            `db:"user_id" json:"-"`
	CreatedAt string         `db:"created_at" json:"-"`
	UpdatedAt string         `db:"updated_at" json:"-"`
	DeletedAt sql.NullString `db:"deleted_at" json:"-"`
	Items     []OrderItems   `json:"items,omitempty"`
}

type OrderItems struct {
	ID        int            `db:"id" json:"-"`
	OrderID   int            `db:"order_id" json:"-"`
	ProductID int            `db:"product_id" json:"-"`
	Product   Product        `json:"product,omitempty"`
	Quantity  int            `db:"quantity" json:"quantity,omitempty"`
	CreatedAt string         `db:"created_at" json:"-"`
	UpdatedAt string         `db:"updated_at" json:"-"`
	DeletedAt sql.NullString `db:"deleted_at" json:"-"`
}

func (o *Orders) FromJSON(data io.Reader) error {
	de := json.NewDecoder(data)
	return de.Decode(o)
}

func (o *Orders) ToJSON() ([]byte, error) {
	return json.Marshal(o)
}
