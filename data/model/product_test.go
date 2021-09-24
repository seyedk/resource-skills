package model

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var productData = `
[
	{
		"id": 1,
		"name": "tesla model y",
		"price": 200.0

	},
	{
		"id": 2,
		"name": "toyota camry LE",
		"price": 1500.0
	}
]
`

func TestProductsDeserriazlizeFromJSON(t *testing.T) {
	p := Products{}

	err := p.FromJSON(bytes.NewReader([]byte(productData)))
	assert.NoError(t, err)

	assert.Len(t, p, 2)
	assert.Equal(t, 1, p[0].ID)
	assert.Equal(t, 2, p[1].ID)
}
func TestProductsSerializesToJSON(t *testing.T) {
	prod := Products{
		Product{ID: 1, Name: "test", Price: 120.12},
	}

	d, err := prod.ToJSON()
	assert.NoError(t, err)

	cd := make([]map[string]interface{}, 0)
	err = json.Unmarshal(d, &cd)
	assert.NoError(t, err)

	assert.Equal(t, float64(1), cd[0]["id"])
	assert.Equal(t, "test", cd[0]["name"])
	assert.Equal(t, float64(120.12), cd[0]["price"])
}
