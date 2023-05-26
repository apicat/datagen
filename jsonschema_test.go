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
					"type":"integer",
					"x-mock":"autoincrement|100,2"
				}
			},
			"createdAt":{
				"type":"integer",
				"x-mock":"timestamp"
			}
		}
	}`

	b, err := JSONSchemaGen([]byte(src), &GenOption{DatagenKey: "x-mock"})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(string(b))

}
