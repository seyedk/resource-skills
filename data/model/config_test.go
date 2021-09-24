package model

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigDeserializeFromJSON( t *testing.T){
	c := Configs{}
	err := c.FromJSON(bytes.NewReader([]byte(configData)))

	assert.NoError(t, err)

	assert.Len(t, c, 3)


	
}

var configData = `
[
	{
		"id": 0,
		"name": "config1",
		"value": "configValue1"
	},
	{
		"id": 1,
		"name": "config2",
		"value": "configValue2"
	},
	{
		"id": 2,
		"name": "config3",
		"value": "configValue2"
	}


]
`
