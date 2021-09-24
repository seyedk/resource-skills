package model

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdersDeserializeFromJSON(t *testing.T) {
	ords := Orders{}
	err := ords.FromJSON(bytes.NewReader([]byte(orderData)))

	assert.NoError(t, err)

}
func TestOrdersSerializeToJSON(t *testing.T){
	ord := Orders{
		Order{
			ID: 1, 
			Items: []OrderItems{
				{
					Product: Product{ID:0},
					Quantity: 1,
				},
				{
					Product: Product{ID:1},
					Quantity: 200,

				},
			},
		},
	}
	d, err := ord.ToJSON()
	assert.NoError(t,err)

	od := make([]map[string]interface{},0)

	err = json.Unmarshal(d,&od)
	assert.NoError(t, err)

	order := od[0]

	assert.Equal(t, float64(1), order["id"])

}

var orderData = `
[
{
	"id":0,
	"items": [
		{
			"product": {
				"id":0,
				"name":"Tesla Model S"
			},
			"quantity":1
		},
		{
			"product": {
				"id":1,
				"name":"Tesla Model Y"
			},
			"quantity":2
		}
	]
},
{
	"id":1,
	"items": [
		{
			"product": {
				"id":0,
				"name":"Tesla Model X"
			},
			"quantity":2
		},
		{
			"product": {
				"id":1,
				"name":"Tesla Model 3"
			},
			"quantity":3
		}
	]
}
]
`
