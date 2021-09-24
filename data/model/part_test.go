package model

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartsDeserializeFromJSON(t *testing.T) {
	products := Products{}
	err := products.FromJSON(bytes.NewReader([]byte(partsData)))

	assert.NoError(t, err)

	assert.Len(t, products, 3)

	assert.Equal(t, 0, products[0].ID)
	assert.Equal(t, 1, products[1].ID)
	assert.Equal(t, 2, products[2].ID)
}

func TestPartsSerializeToJSON(t *testing.T) {
	parts := Parts{
		Part{ID: 1, Name: "test", Quantity: 10, Unit: "cnt"},
	}

	j, err := parts.ToJSON()

	assert.NoError(t, err)
	newParts := make([]map[string]interface{}, 0)

	err = json.Unmarshal(j, &newParts)

	assert.NoError(t, err)
	
	assert.Equal(t, float64(1), newParts[0]["id"])
	assert.Equal(t, "test", newParts[0]["name"])
	assert.Equal(t, float64(10), newParts[0]["quantity"])
	

}

var partsData = `
[
	{
		"id": 0,
		"name": "Auto Pilot",
		"quantity": 40,
		"unit": "cnt"
	},
	{
		"id": 1,
		"name": "Self Driving",
		"quantity": 3,
		"unit": "cnt"
	},
	{
		"id": 2,
		"name": "Alloy Wheels",
		"quantity": 2,
		"unit": "cnt"
	}
]
`
