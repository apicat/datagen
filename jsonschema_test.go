package datagen

import (
	"fmt"
	"testing"
)

func TestJSONSchema(t *testing.T) {
	src := `{
		"type":"object",
		"properties":{
			"id":{
				"type":"string",
				"format":"uuid"
			},
			"name":{
				"type":"string",
				"x-mock":"name"
			},
			"address":{
				"type":"string",
				"x-mock":"address"
			},
			"books":{
				"type":"array",
				"items":{
					"type":"number"
				}
			},
			"createdAt":{
				"type":"integer",
				"x-mock":"timestamp"
			}
		}
	}`

	b, err := JSONSchemaGen([]byte(src), &JSONSchemaOption{})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(string(b))

}